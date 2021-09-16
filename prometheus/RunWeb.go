package prometheus

import (
	logg "log"
	"net/http"

	"github.com/myeongsuk.yoon/metrics-operator/prometheus/metricsRouter"
	"github.com/prometheus/common/log"
)

func RunWebServer() {
	logg.Println("Metrics Web Start")

	err := http.ListenAndServe(":8082", metricsRouter.MetricsRouter)
	if err != nil {
		log.Errorln(err)
	}
	logg.Println("Metrics Web End")
}
