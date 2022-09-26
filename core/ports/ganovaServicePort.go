package ports

import (
	model "github.com/jearcila/hex-architecture/core/model"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
)

type GenovaServiceInt interface {
	Authorize(tctx transactions_context.Context, authorization model.FirstOperationRequest) (model.FirstOperationResponse, error)
	Capture(tctx transactions_context.Context, capture model.CaptureRequest) (model.CaptureResponse, error)
	Cancel(tctx transactions_context.Context, cancel model.CancelRequest) (model.CancelResponse, error)
	Purchase(tctx transactions_context.Context, purchase model.FirstOperationRequest) (model.FirstOperationResponse, error)
}
