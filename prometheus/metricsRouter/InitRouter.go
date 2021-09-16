package metricsRouter

import (
	"github.com/myeongsuk.yoon/metrics-operator/prometheus/scrapeMetrics"
)

var MetricsRouter *router

func init() {
	MetricsRouter = new(router)
	MetricsRouter.targetMap = make(map[string]scrapeMetrics.TargetGroup)
}
