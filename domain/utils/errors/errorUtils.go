package errors

import (
	"encoding/json"
	"errors"

	"github.com/jearcila/hex-architecture/domain/constants/status"
	gErrors "github.com/jearcila/hex-architecture/domain/model/errors"
	utils "github.com/jearcila/hex-architecture/domain/utils"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
	transactions_models "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/models"
)

func ErrorResponse(context transactions_context.Context, elapsed *float64, err error) interface{} {
	var errorResponse *gErrors.ErrorResponse
	if !errors.As(err, &errorResponse) {
		return transactions_factory.CONTINGENCY(context)
	}
	var operationStatus transactions_factory.OperationStatus
	var exist bool

	operationStatus, exist = status.StatusByProviderAuthorization[utils.GetStatusTag(context, errorResponse.ErrorCode)]
	if !exist {
		operationStatus = utils.SetDefaultOperationStatus(context, errorResponse.ErrorCode)
	}

	operationStatus.Descriptor = utils.GetDescriptor(context)

	operationStatus.Elapsed = elapsed
	operationStatus.ProviderStatusCode = errorResponse.ErrorCode
	operationStatus.ProviderStatus = errorResponse.ErrorMessage
	operationStatus.Merchant = context.Transaction.Merchant.Account
	operationStatus.Reference = transactions_models.Reference{}
	operationStatus.Response = utils.GetRawResponseAsString(context, errorResponse)

	rawProviderJSON := json.RawMessage(operationStatus.Response)
	operationStatus.Parsed = rawProviderJSON

	resp, err := transactions_factory.TransactionResponse(context, operationStatus)
	if err != nil {
		return err
	}
	return resp
}
