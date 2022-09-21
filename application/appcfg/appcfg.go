package appcfg

import (
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/furyconfig"
)

func GetString(context transactions_context.Context, key string) string {
	return context.FuryConfig.Get(key, "", func(value string) error { return nil })
}

func GetStringFromFuryConfig(furyConfig furyconfig.Config, key string) string {
	return furyConfig.Get(key, "", func(value string) error { return nil })
}
