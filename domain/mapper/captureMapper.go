package mapper

import (
	"encoding/json"

	"github.com/jearcila/hex-architecture/domain/constants/status"
	"github.com/jearcila/hex-architecture/domain/model"
	utils "github.com/jearcila/hex-architecture/domain/utils"
	log "github.com/jearcila/hex-architecture/domain/utils/log"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
)

func CaptureTransactionResponse(context transactions_context.Context, elapsed *float64, captureResponse model.CaptureResponse) (interface{}, error) {
	operationStatus := status.FindStatusByResponseCode(status.StatusByProviderCapture, captureResponse.ResponseCode)

	operationStatus.Descriptor = utils.GetDescriptor(context)

	operationStatus.Elapsed = elapsed
	operationStatus.ProviderStatusCode = captureResponse.ResponseCode
	operationStatus.ProviderStatus = captureResponse.ResponseMessage
	operationStatus.Merchant = context.Transaction.Merchant.Account
	operationStatus.AuthorizationCode = captureResponse.AuthorizationCode

	references, err := utils.SaveSecondOperationReferences(context)
	if err != nil {
		log.EventError(context, log.EventErrorNotSavedReferences, err)
	}

	operationStatus.Reference = references
	operationStatus.Response = utils.GetRawResponseAsString(context, captureResponse)

	rawProviderJSON := json.RawMessage(operationStatus.Response)
	operationStatus.Parsed = rawProviderJSON

	return transactions_factory.TransactionResponse(context, operationStatus)
}
