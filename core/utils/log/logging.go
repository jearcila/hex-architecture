package log

import (
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/logger"
)

// LogEvent logs the event as info level
func Event(context transactions_context.Context, event string) {
	context.Log(logger.INFO, event)
}

// EventInfoWithError logs the event and notice the error
func EventInfoWithError(context transactions_context.Context, event string, err error) {
	tags := logger.Tags{
		TagErrorMessage: err.Error(),
	}
	context.Log(logger.INFO, event, tags)
}

// LogEventError logs the event and the error
func EventError(context transactions_context.Context, event string, err error) {
	tags := logger.Tags{
		TagErrorMessage: err.Error(),
	}
	context.Log(logger.ALERT, event, tags)
}

// LogBuildData logs generated data info
func RawMessage(context transactions_context.Context, event string, data string) {
	tags := logger.Tags{
		TagRawMessage: data,
	}

	context.Log(logger.INFO, event, tags)
}

func BrandNotDeterminateInfo(context transactions_context.Context, msg string) {
	tags := logger.Tags{
		"message": msg,
	}
	context.Log(logger.INFO, EventBrandNotDeterminate, tags)
}
