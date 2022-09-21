package mapper

import (
	"encoding/json"

	"github.com/jearcila/hex-architecture/domain/constants/status"
	"github.com/jearcila/hex-architecture/domain/model"
	"github.com/jearcila/hex-architecture/domain/utils"
	log "github.com/jearcila/hex-architecture/domain/utils/log"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
	mapping "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/mapping"
)

func PurchaseTransactionResponse(context transactions_context.Context, elapsed *float64, purchaseResponse model.FirstOperationResponse) (interface{}, error) {
	operationStatus := status.FindStatusByResponseCode(status.StatusByProviderPurchase, purchaseResponse.ResponseCode)

	traditionalStatus, isTraditionalResponse := mapping.Find(context, []mapping.Mapping{
		mapping.TraditionalCardsMapping(purchaseResponse.ResponseCode, operationStatus.Status),
	})

	if isTraditionalResponse {
		operationStatus.Status = traditionalStatus
	}

	operationStatus.Descriptor = utils.GetDescriptor(context)

	operationStatus.Elapsed = elapsed
	operationStatus.ProviderStatusCode = purchaseResponse.ResponseCode
	operationStatus.ProviderStatus = purchaseResponse.ResponseMessage
	operationStatus.Merchant = context.Transaction.Merchant.Account
	operationStatus.AuthorizationCode = purchaseResponse.AuthorizationCode

	references, err := utils.SaveAuthorizationReferences(context, purchaseResponse)
	if err != nil {
		log.EventError(context, log.EventErrorNotSavedReferences, err)
	}

	operationStatus.Reference = references
	operationStatus.Response = utils.GetRawResponseAsString(context, purchaseResponse)

	rawProviderJSON := json.RawMessage(operationStatus.Response)
	operationStatus.Parsed = rawProviderJSON

	return transactions_factory.TransactionResponse(context, operationStatus)
}
