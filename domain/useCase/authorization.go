package useCase

import (
	"time"

	constants "github.com/jearcila/hex-architecture/domain/constants"
	"github.com/jearcila/hex-architecture/domain/mapper"
	"github.com/jearcila/hex-architecture/domain/model"
	ports "github.com/jearcila/hex-architecture/domain/ports"
	"github.com/jearcila/hex-architecture/domain/utils"
	"github.com/jearcila/hex-architecture/domain/utils/errors"
	"github.com/jearcila/hex-architecture/domain/utils/format"
	"github.com/jearcila/hex-architecture/domain/utils/log"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
	transactions_models "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/models"
)

type G2Handler func(context transactions_context.Context, transaction transactions_models.Transaction) interface{}

func Authorization(acquirer ports.GenovaServiceInt) G2Handler {
	return func(ctx transactions_context.Context, transaction transactions_models.Transaction) interface{} {
		// Only allows operations for genova provider
		if ctx.Transaction.Provider.Id != constants.GenovaIntegration {
			return transactions_factory.INCORRECT_INTEGRATION_RECEIVER(ctx, ctx.Transaction.Provider.Id)
		}

		log.Event(ctx, log.EventBuildAuthorizationMessage)
		authorizationRequest, err := buildAuthorization(ctx)
		if err != nil {
			log.EventError(ctx, log.EventErrorBuildAuthorizationMessage, err)
			return transactions_factory.ERROR_BUILDING_CHANNEL_DATA(ctx, err.Error())
		}

		// Send transaction to Genova
		message, err := format.GetRawJSONAsString(authorizationRequest)
		if err != nil {
			log.EventError(ctx, log.EventErrorBuildAuthorizationMessage, err)
			return transactions_factory.ERROR_BUILDING_CHANNEL_DATA(ctx, err.Error())
		}
		log.RawMessage(ctx, log.EventSendAuthorizationMessage, message)

		timeBeforeProcess := time.Now()
		authorizationResponse, err := acquirer.Authorize(ctx, authorizationRequest)
		totalDuration := time.Since(timeBeforeProcess).Seconds()
		if err != nil {
			log.EventInfoWithError(ctx, log.EventErrorAcquirerResponse, err)
			return errors.ErrorResponse(ctx, &totalDuration, err)
		}

		response, err := format.GetRawJSONAsString(authorizationResponse)
		if err != nil {
			log.EventInfoWithError(ctx, log.EventErrorAcquirerResponse, err)
			return errors.ErrorResponse(ctx, &totalDuration, err)
		}
		log.RawMessage(ctx, log.EventAuthorizationResponse, response)

		// Create transaction response
		log.Event(ctx, log.EventCreateAuthorizationResponse)
		resp, err := mapper.AuthorizationTransactionResponse(ctx, &totalDuration, authorizationResponse)
		if err != nil {
			log.EventError(ctx, log.EventErrorCreateAuthorizationResponse, err)
			return err
		}
		return resp
	}
}

func buildAuthorization(ctx transactions_context.Context) (model.FirstOperationRequest, error) {
	// Validate data to fill operation structure
	var req model.FirstOperationRequest
	if err := utils.ParseContext(ctx, &req); err != nil {
		return model.FirstOperationRequest{}, err
	}

	return req, nil
}
