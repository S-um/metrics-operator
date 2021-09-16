package scrapeMetrics

import (
	"fmt"
	"log"
	"net/http"
)

type TargetConfig interface {
	scrapeMetrics() (string, error)
	GetName() string
	GetPath() string
	GetProtocol() string
	GetPort() string
}

type TargetGroup struct {
	name    string
	configs []TargetConfig
}

func NewTargetGroup(name string, configs []TargetConfig) TargetGroup {
	newGroup := TargetGroup{}
	newGroup.name = name
	newGroup.configs = configs

	return newGroup
}

func (t *TargetGroup) ServeFunc(w http.ResponseWriter, r *http.Request) {
	if t.configs != nil {
		for _, config := range t.configs {
			metrics, err := config.scrapeMetrics()
			if err != nil {
				log.Println("Fail to scrape service from", t.name, ":", config.GetName()+"/"+config.GetPath())
				log.Println(err)
			} else {
				log.Println("Success to scrape service from", t.name, ":", config.GetName()+"/"+config.GetPath())
				fmt.Fprint(w, metrics)
			}
		}
	}
}
