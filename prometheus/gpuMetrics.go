package prometheus

import (
	"bytes"
	"fmt"
	logg "log"
	"net/http"

	"github.com/prometheus/common/log"
)

func custom(w http.ResponseWriter, r *http.Request) {
	logg.Println("Custom Request Start")
	fmt.Fprint(w, "custom custom")
	w.WriteHeader(http.StatusOK)
	logg.Println("Custom Request End")
}

func gpuMetrics(w http.ResponseWriter, r *http.Request) {
	logg.Println("Gpu Request Start")
	resp, err := http.Get("http://dcgm-exporter:9400/metrics")
	if err != nil {
		log.Errorln(err)
		fmt.Fprintf(w, "no response")
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Fprint(w, buf.String())
	fmt.Fprint(w, "hello, it's me!!")
	fmt.Print(buf.String())
	w.WriteHeader(http.StatusOK)
	logg.Println("Gpu Request End")
}
