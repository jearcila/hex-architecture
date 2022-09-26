package mapper

import (
	"encoding/json"

	"github.com/jearcila/hex-architecture/core/constants/status"
	model "github.com/jearcila/hex-architecture/core/model"
	utils "github.com/jearcila/hex-architecture/core/utils"
	log "github.com/jearcila/hex-architecture/core/utils/log"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
	mapping "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/mapping"
)

func AuthorizationTransactionResponse(context transactions_context.Context, elapsed *float64, response model.FirstOperationResponse) (interface{}, error) {

	var operationStatus transactions_factory.OperationStatus
	var exist bool

	hybridStatus, isHybridResponse := mapping.Find(context, []mapping.Mapping{
		mapping.HybridCardsMapping(response.ResponseCode),
	})

	if isHybridResponse {
		operationStatus = transactions_factory.OperationStatus{
			Status: hybridStatus,
		}
	} else {
		operationStatus, exist = status.StatusByProviderAuthorization[utils.GetStatusTag(context, response.ResponseCode)]
		if !exist {
			operationStatus = utils.SetDefaultOperationStatus(context, response.ResponseCode)
		} else {
			utils.ValidateExceedRetryLimit(&context, &operationStatus)
		}
	}

	traditionalStatus, isTraditionalResponse := mapping.Find(context, []mapping.Mapping{
		mapping.TraditionalCardsMapping(response.ResponseCode, operationStatus.Status),
	})

	if isTraditionalResponse {
		operationStatus.Status = traditionalStatus
	}

	operationStatus.Descriptor = utils.GetDescriptor(context)

	operationStatus.Elapsed = elapsed
	operationStatus.ProviderStatusCode = response.ResponseCode
	operationStatus.ProviderStatus = response.ResponseMessage
	operationStatus.Merchant = context.Transaction.Merchant.Account
	operationStatus.AuthorizationCode = response.AuthorizationCode

	references, err := utils.SaveAuthorizationReferences(context, response)
	if err != nil {
		log.EventError(context, log.EventErrorNotSavedReferences, err)
	}

	operationStatus.Reference = references
	operationStatus.Response = utils.GetRawResponseAsString(context, response)

	rawProviderJSON := json.RawMessage(operationStatus.Response)
	operationStatus.Parsed = rawProviderJSON

	return transactions_factory.TransactionResponse(context, operationStatus)
}
