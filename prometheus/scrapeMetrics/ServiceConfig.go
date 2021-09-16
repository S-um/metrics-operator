package scrapeMetrics

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type ServiceConfig struct {
	Protocol string `json:"protocol,omitempty"`
	Name     string `json:"name"`
	Path     string `json:"path,omitempty"`
	Port     string `json:"port,omitempty"`
}

func ServiceArrToTargetConfigArr(s []ServiceConfig) []TargetConfig {
	length := len(s)
	arr := make([]TargetConfig, length)
	for index := 0; index < length; index++ {
		arr[index] = &s[index]
	}
	return arr
}

func (s *ServiceConfig) scrapeMetrics() (string, error) {
	endpoints, err := s.getEndpoint()
	if err != nil {
		return "", err
	}
	metrics := ""
	for _, ep := range endpoints {
		log.Println("Get metrics from", s.Protocol+"://"+ep+":"+s.Port+"/"+s.Path)
		resp, err := http.Get(s.Protocol + "://" + ep + ":" + s.Port + "/" + s.Path)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		str, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		} else {
			metrics += string(str)
		}
	}

	return metrics, nil
}

func (s *ServiceConfig) getEndpoint() ([]string, error) {
	endpoints := []string{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//https://223.194.90.117:8443/api/v1/namespaces/default/endpoints/dcgm-exporter
	log.Println("Get endpoints from", "https://"+os.Getenv("CLUSTERIP")+"/api/v1/namespaces/metrics-operator-system/endpoints/"+s.Name)
	resp, err := client.Get("https://" + os.Getenv("CLUSTERIP") + "/api/v1/namespaces/metrics-operator-system/endpoints/" + s.Name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	temp := strings.Split(string(data), "\"ip\":")
	for i := 1; i < len(temp); i++ {
		temp2 := strings.Split(temp[i], ",")
		endpoints = append(endpoints, strings.Replace(temp2[0], "\"", "", -1))
	}

	return endpoints, nil
}

func (s *ServiceConfig) GetName() string {
	return s.Name
}

func (s *ServiceConfig) GetPath() string {
	return s.Path
}

func (s *ServiceConfig) GetPort() string {
	return s.Port
}

func (s *ServiceConfig) GetProtocol() string {
	return s.Protocol
}
