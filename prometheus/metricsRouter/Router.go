package metricsRouter

import (
	"fmt"
	"net/http"

	v1 "github.com/myeongsuk.yoon/metrics-operator/api/v1"
	"github.com/myeongsuk.yoon/metrics-operator/prometheus/scrapeMetrics"
)

type router struct {
	targetMap map[string]scrapeMetrics.TargetGroup
}

func UpdateConfig(target v1.MetricsTarget, name string) {
	MetricsRouter.updateConfig(target, name)
}

func (t *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serve, exist := t.targetMap[r.URL.String()]
	if exist {
		w.WriteHeader(http.StatusOK)
		serve.ServeFunc(w, r)
		return
	}
	if r.URL.String() == PathPrefix+PathSuffix {
		w.WriteHeader(http.StatusOK)
		for _, targetGroup := range t.targetMap {
			targetGroup.ServeFunc(w, r)
		}
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "wrong url path request")
	fmt.Fprintln(w, PathPrefix+PathSuffix)
	for key, _ := range t.targetMap {
		fmt.Fprintln(w, key)
	}
}
