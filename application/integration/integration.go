package integration

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	appcfg "github.com/jearcila/hex-architecture/application/appcfg"
	enviroment "github.com/jearcila/hex-architecture/application/enviroments"
	genova "github.com/jearcila/hex-architecture/infraestructure/driven-adapters/services"
	apirest "github.com/jearcila/hex-architecture/infraestructure/entry-points/api-rest"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/integrations"
	transactions_constants "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/constants"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/furyconfig"
	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-core/pkg/telemetry"
	"github.com/mercadolibre/fury_go-core/pkg/telemetry/tracing"
	"github.com/mercadolibre/fury_go-core/pkg/transport/httpclient"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const _integrationName = "genova"

func Run() error {
	// Init web app
	app, err := fury.NewWebApplication(fury.WithLogLevel(log.DebugLevel))
	if err != nil {
		return fmt.Errorf("could not init web app: %w", err)
	}

	integration := SetupIntegration()

	regularAuthorization := integrations.Operation{
		Type: transactions_constants.AUTHORIZATION,
		Mode: transactions_constants.MODE_REGULAR,
	}

	regularCapture := integrations.Operation{
		Type: transactions_constants.CAPTURE,
		Mode: transactions_constants.MODE_REGULAR,
	}

	regularRefund := integrations.Operation{
		Type: transactions_constants.REFUND,
		Mode: transactions_constants.MODE_REGULAR,
	}

	regularPurchase := integrations.Operation{
		Type: transactions_constants.PURCHASE,
		Mode: transactions_constants.MODE_REGULAR,
	}

	// Set this value to get the correct environment name
	app.Scope.Environment = enviroment.GetEnv(app.Scope.Environment)

	// Init genova service
	genovaService, err := NewService(enviroment.Environment(app.Scope.Environment))
	if err != nil {
		return fmt.Errorf("error creating genova service: %w", err)
	}
	// Crear instancia del servicio rest
	api := apirest.CreateInstance()
	// register telemetry middleware
	telemetryMiddleware := func(engine *gin.Engine) {
		engine.Use(telemetryMiddleware(app.Tracer, app.Logger))
	}

	// Register Integration operations
	integration.OnlineOperation(regularAuthorization, integrations.OnlineOperationHandler(api.AuthorizationRest(&genovaService)))
	integration.OnlineOperation(regularCapture, integrations.OnlineOperationHandler(api.CaptureRest(&genovaService)))
	integration.OnlineOperation(regularRefund, integrations.OnlineOperationHandler(api.PurchaseRest(&genovaService)))
	integration.OnlineOperation(regularPurchase, integrations.OnlineOperationHandler(api.RefundRest(&genovaService)))

	integration.Deploy(telemetryMiddleware)
	return nil

}

func SetupIntegration() *integrations.Integration {
	mandatoryConfigs := []string{
		appcfg.GenovaURLLocalFuryConfigKey,
		appcfg.GenovaURLProdFuryConfigKey,
		appcfg.GenovaURLStagingFuryConfigKey,
		appcfg.GenovaProcessorID,
	}

	integration, integrationError := integrations.New(_integrationName,
		integrations.WithFuryConfig(furyconfig.New(_integrationName)),
		integrations.WithMandatoryFuryConfigs(mandatoryConfigs),
		integrations.WithSupportedCurrencies([]string{"BRL", "ARS", "MXN", "CLP"}),
		integrations.WithoutValidation(integrations.INTEGRATION_NAME_MUST_MATCH_TRANSACTION_PROVIDER_ID),
		integrations.WithCapabilities(GetCapabilities()),
	)

	if integrationError != nil {
		panic(fmt.Sprintf("Could not initialize Integration: %s", integrationError.Error()))
	}

	return integration
}

func telemetryMiddleware(tracer telemetry.Client, logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// G2 integration only injects NewRelic context into gin's context, we
		// want it to be present in c.Request.Context() as well if not present.
		if newrelic.FromContext(ctx) == nil {
			tx := nrgin.Transaction(c)
			if tx != nil {
				ctx = newrelic.NewContext(ctx, tx)
			}
		}

		// Decorate context with tracing information we are required to propagate.
		ctx = tracing.ContextFromHeader(ctx, c.Request.Header)

		// The scope of the variable logger is that of the parent function, which
		// on most programs will be called only once. In contrast, this function
		// will be called on a per transaction basis, meaning that assignment of the
		// logger variable directly causes a race condition and unexpected
		// behavior. To avoid that we assign the logger to a variable of local
		// scope, which allows us to change it without side effects.
		l := logger

		// Check if transaction-id is present, and add it to the logger then
		requestID := c.Request.Header.Get(tracing.RequestIDHeader)
		if requestID != "" {
			l = l.With(log.String("request_id", requestID))
		}

		// Add telemetry client on context
		ctx = telemetry.Context(ctx, tracer)

		// Add logger client on context
		ctx = log.Context(ctx, l)

		// Recreate transaction with the decorated context
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func NewService(env enviroment.Environment) (genova.Client, error) {
	requester := httpclient.New(httpclient.WithTimeout(240 * time.Second))
	service, err := genova.NewService(env, requester)
	if err != nil {
		return genova.Client{}, err
	}

	return genova.Client{
		Service: service,
	}, nil
}
