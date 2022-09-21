package genovaService

import (
	"encoding/json"
	"net/http"

	appcfg "github.com/jearcila/hex-architecture/application/appcfg"
	environment "github.com/jearcila/hex-architecture/application/enviroments"
	errors "github.com/jearcila/hex-architecture/domain/model/errors"
	apimodel "github.com/mercadolibre/fury_gateway-kit/pkg/g2/journal/client/model"

	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/furyconfig"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/journal"
	"github.com/mercadolibre/fury_go-core/pkg/rusty"
	"github.com/mercadolibre/fury_go-core/pkg/transport/httpclient"
)

var _environmentBaseURL = map[environment.Environment]string{
	environment.EnvironmentLocal:      appcfg.GenovaURLLocalFuryConfigKey,
	environment.EnvironmentProduction: appcfg.GenovaURLProdFuryConfigKey,
	environment.EnvironmentStaging:    appcfg.GenovaURLStagingFuryConfigKey,
}

const (
	_methodAuthorize = "/authorize"
	_methodCapture   = "/capture"
	_methodCancel    = "/cancel"
	_methodPurchase  = "/purchase"
)

type service struct {
	endpointAuthorize *rusty.Endpoint
	endpointCapture   *rusty.Endpoint
	endpointCancel    *rusty.Endpoint
	endpointPurchase  *rusty.Endpoint
}

func NewService(env environment.Environment, requester httpclient.Requester) (*service, error) {
	urlFuryConfigKey, exists := _environmentBaseURL[env]
	if !exists {
		return nil, errors.ErrUnknownEnvironment
	}

	genovaFuryConfig := furyconfig.New("genova")
	url := appcfg.GetStringFromFuryConfig(genovaFuryConfig, urlFuryConfigKey)

	var errorPolicy rusty.ErrorPolicyFunc = func(r *rusty.Response) error {
		if r.StatusCode < http.StatusBadRequest {
			return nil
		}
		var errorResponse errors.ErrorResponse
		if err := json.Unmarshal(r.Body, &errorResponse); err != nil {
			return err
		}
		switch errorResponse.ErrorCode {
		case errors.ErrorCodeE001, errors.ErrorCodeE002, errors.ErrorCodeE003, errors.ErrorCodeE004, errors.ErrorCodeE005:
			return errors.NewErrorResponse(errorResponse.ErrorCode, errorResponse.ErrorMessage)
		}
		return errors.NewUnexpectedClientResponse(r.StatusCode, r.Body)

	}
	headers := rusty.WithHeader("content-type", "application/json")

	params := []rusty.EndpointOption{
		headers,
		rusty.WithErrorPolicy(errorPolicy),
	}
	endpointAuthorize, err := rusty.NewEndpoint(requester,
		rusty.URL(url, _methodAuthorize),
		params...)
	if err != nil {
		return nil, err
	}

	endpointCapture, err := rusty.NewEndpoint(requester,
		rusty.URL(url, _methodCapture),
		params...)
	if err != nil {
		return nil, err
	}

	endpointCancel, err := rusty.NewEndpoint(requester,
		rusty.URL(url, _methodCancel),
		params...)
	if err != nil {
		return nil, err
	}

	endpointPurchase, err := rusty.NewEndpoint(requester,
		rusty.URL(url, _methodPurchase),
		params...)
	if err != nil {
		return nil, err
	}

	return &service{
		endpointAuthorize: endpointAuthorize,
		endpointCapture:   endpointCapture,
		endpointCancel:    endpointCancel,
		endpointPurchase:  endpointPurchase,
	}, nil
}

func saveRequestInJournal(tctx transactions_context.Context, request interface{}) {
	event := apimodel.Event{
		Rating: "INFO",
		ID:     "authorization_router_request",
		Entity: tctx.Transaction.Id,
		Component: apimodel.Component{
			ID:   tctx.Transaction.Provider.Id,
			Type: tctx.Layer,
		},
		Info: request,
		MetaData: apimodel.MetaData{
			HTTP: apimodel.MetaDataHTTP{
				XRequestId: tctx.RequestId,
			},
			Duplicates: apimodel.MetaDataDuplicates{
				OperationAttemptID: tctx.OperationAttemptID,
			},
		},
	}
	journal.Save(event)
}

func saveResponseInJournal(tctx transactions_context.Context, response *rusty.Response) {
	event := apimodel.Event{
		Rating: "INFO",
		ID:     "authorization_router_response",
		Entity: tctx.Transaction.Id,
		Component: apimodel.Component{
			ID:   tctx.Transaction.Provider.Id,
			Type: tctx.Layer,
		},
		Info: json.RawMessage(response.Body),
		MetaData: apimodel.MetaData{
			HTTP: apimodel.MetaDataHTTP{
				XRequestId: tctx.RequestId,
			},
			Duplicates: apimodel.MetaDataDuplicates{
				OperationAttemptID: tctx.OperationAttemptID,
			},
		},
	}
	journal.Save(event)
}

func saveErrorInJournal(tctx transactions_context.Context, err error) {
	event := apimodel.Event{
		Rating: "INFO",
		ID:     "authorization_router_error",
		Entity: tctx.Transaction.Id,
		Component: apimodel.Component{
			ID:   tctx.Transaction.Provider.Id,
			Type: tctx.Layer,
		},
		Info: err.Error(),
		MetaData: apimodel.MetaData{
			HTTP: apimodel.MetaDataHTTP{
				XRequestId: tctx.RequestId,
			},
			Duplicates: apimodel.MetaDataDuplicates{
				OperationAttemptID: tctx.OperationAttemptID,
			},
		},
	}
	journal.Save(event)
}
