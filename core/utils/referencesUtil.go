package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jearcila/hex-architecture/core/constants"
	"github.com/jearcila/hex-architecture/core/model"
	transactions_constants "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/constants"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
	transactions_models "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/models"
)

func SaveAuthorizationReferences(ctx transactions_context.Context, authorizationResponse model.FirstOperationResponse) (transactions_models.Reference, error) {
	if authorizationResponse.AcquirerTransactionID == "" {
		return transactions_models.Reference{}, errors.New("not found acquirer transaction id")
	}

	merchantTransactionReference := authorizationResponse.AcquirerTransactionID
	ref := transactions_factory.BuildResponseReference(merchantTransactionReference)

	if authorizationResponse.ICCRelatedData != "" {
		transactions_factory.AddReference(ref, constants.IccRelatedData, authorizationResponse.ICCRelatedData)
	}

	if authorizationResponse.ResponseCode != "" {
		responseCode := authorizationResponse.ResponseCode

		// For compatibility reasons with point, if code is an internal error (EXXX), we map it to ISO code 05
		if strings.HasPrefix(responseCode, "E") {
			responseCode = "05" // Not using constant as import cycle fix would make things more complex than needed
		}

		transactions_factory.AddReference(ref, constants.IsoResponseCode, responseCode)
	}

	transactions_factory.AddReference(ref, constants.AcquirerTransactionID, merchantTransactionReference)

	transactions_factory.AddReference(ref, constants.MerchantOperationReference, BuildMerchantOperationReference(ctx))

	transactions_factory.AddReference(ref, constants.ReconciliationTicketID, BuildMerchantOperationReference(ctx))
	return ref, nil
}

func SaveSecondOperationReferences(ctx transactions_context.Context) (transactions_models.Reference, error) {
	merchantTransactionReference, err := GetAcquirerTransactionID(ctx)
	if err != nil {
		return nil, err
	}

	ref := transactions_factory.BuildResponseReference(merchantTransactionReference)
	transactions_factory.AddReference(ref, constants.AcquirerTransactionID, merchantTransactionReference)
	transactions_factory.AddReference(ref, constants.MerchantOperationReference, BuildMerchantOperationReference(ctx))
	transactions_factory.AddReference(ref, constants.ReconciliationTicketID, BuildMerchantOperationReference(ctx))
	return ref, nil
}

func BuildMerchantOperationReference(ctx transactions_context.Context) string {
	splitID := ctx.Transaction.Id
	var operationSuffix string

	switch ctx.Transaction.Operation.Type {
	case transactions_constants.AUTHORIZATION:
		operationSuffix = "A"
	case transactions_constants.CAPTURE:
		operationSuffix = "C"
	case transactions_constants.PURCHASE:
		operationSuffix = "P"
	case transactions_constants.REFUND:
		refundID := strconv.FormatUint(ctx.Transaction.Operation.RefundId, 10)
		operationSuffix = "R" + "_" + refundID
	}
	return splitID + "_" + operationSuffix
}

func GetAcquirerTransactionID(ctx transactions_context.Context) (string, error) {
	references := GetFirstOperationReferences(*ctx.Transaction)
	acquirerTransactionID, ok := references[constants.AcquirerTransactionID]
	if !ok {
		return "", errors.New("acquirer_transaction_id reference not found")
	}
	return acquirerTransactionID, nil
}

func GetFirstOperationReferences(transaction transactions_models.Transaction) map[string]string {
	if transaction.Operation.References == nil {
		return map[string]string{}
	}
	if transaction.Operation.References.Authorization != nil && len(transaction.Operation.References.Authorization.Reference) != 0 {
		return transaction.Operation.References.Authorization.Reference
	}
	if transaction.Operation.References.Purchase != nil && len(transaction.Operation.References.Purchase.Reference) != 0 {
		return transaction.Operation.References.Purchase.Reference
	}

	return map[string]string{}
}
