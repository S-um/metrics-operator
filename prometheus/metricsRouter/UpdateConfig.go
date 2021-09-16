package metricsRouter

import (
	"fmt"

	v1 "github.com/myeongsuk.yoon/metrics-operator/api/v1"
	"github.com/myeongsuk.yoon/metrics-operator/prometheus/scrapeMetrics"
)

func checkActiveConfig(target v1.MetricsTargetSpec) bool {
	if target.Service == nil && target.Url == nil {
		return false
	}
	return true
}

func (r *router) updateConfig(target v1.MetricsTarget, name string) {
	if checkActiveConfig(target.Spec) == true {
		serviceTarget := []scrapeMetrics.TargetConfig(scrapeMetrics.ServiceArrToTargetConfigArr(target.Spec.Service))
		urlTarget := []scrapeMetrics.TargetConfig(scrapeMetrics.UrlArrToTargetConfigArr(target.Spec.Url))

		targetGroup := scrapeMetrics.NewTargetGroup(target.Name, append(serviceTarget, urlTarget...))
		r.targetMap[PathPrefix+"/"+target.Name+PathSuffix] = targetGroup
	} else {
		fmt.Println("Delete :", PathPrefix+"/"+name+PathSuffix)
		delete(r.targetMap, PathPrefix+"/"+name+PathSuffix)
	}
}
