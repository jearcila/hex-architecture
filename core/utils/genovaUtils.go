package utils

import (
	"fmt"
	"strings"

	"github.com/jearcila/hex-architecture/core/constants/status"
	"github.com/jearcila/hex-architecture/core/utils/format"
	"github.com/jearcila/hex-architecture/core/utils/log"
	"github.com/jearcila/hex-architecture/core/utils/metrics"

	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_factory "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/factory"
	transactions_models "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/models"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/descriptor"
)

func GetRawResponseAsString(context transactions_context.Context, errorResponse interface{}) string {
	rawResponse, err := format.GetRawJSONAsString(errorResponse)
	if err != nil {
		log.EventError(context, log.EventErrorNotSavedResponse, err)
	}
	return rawResponse
}

func GetStatusTag(ctx transactions_context.Context, respCode string) string {
	if ctx.Transaction == nil || ctx.Transaction.Operation.Card == nil || ctx.Transaction.Operation.Card.Brand == "" {
		log.BrandNotDeterminateInfo(ctx, log.MessageErrorBrandNoDeterminate)
		return ""
	}
	brand := GetBrandKey(ctx.Transaction.Operation.Card.Brand)
	if ctx.Transaction.Operation.Card.Present != nil {
		return fmt.Sprintf("%s%s%s", respCode, brand, status.CardPresent)
	}
	return fmt.Sprintf("%s%s%s", respCode, brand, status.Ecommerce)
}

func GetBrandKey(brand string) string {
	defa := ""
	if strings.Contains(brand, "master") {
		defa = status.Master
	}
	if strings.Contains(brand, "visa") {
		defa = status.Visa
	}
	return defa
}

func SetDefaultOperationStatus(ctx transactions_context.Context, respCode string) transactions_factory.OperationStatus {
	if ctx.Transaction.Operation.Card.Present != nil {
		return status.FindStatusByResponseCode(status.StatusByProviderAuthorizationCardPresent, respCode)
	}
	return status.FindStatusByResponseCode(status.StatusByProviderAuthorizationEcommerce, respCode)
}

func ValidateExceedRetryLimit(ctx *transactions_context.Context, opStatus *transactions_factory.OperationStatus) {
	limit, isPresent := status.LimitRetryNumber[ctx.Transaction.Operation.Card.Brand]
	if opStatus.Status == status.Contingency.Status && isPresent && ctx.Transaction.Operation.Retries >= limit {
		*opStatus = status.Rejected
		metrics.MetricForAbecsRetriesLimit(ctx.Transaction.Operation.Card.Brand)
	}
}

func GetDescriptor(context transactions_context.Context) *transactions_models.DescriptorResponse {
	desc := descriptor.BuildDescriptorWithMerchantAndProduct(context, descriptor.Limit{}, "")
	return &desc
}
