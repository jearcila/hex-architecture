package metrics

import (
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/version"
	"github.com/mercadolibre/go-meli-toolkit/godog"
)

func MetricForAbecsRetriesLimit(brand string) {
	metric := "application.g2.abecs.retries.limit"

	tags := new(godog.Tags)
	tags.Add("provider", "genova")
	tags.Add("version", version.Version())
	tags.Add("brand", brand)

	godog.RecordSimpleMetric(metric, 1, tags.ToArray()...)
}
