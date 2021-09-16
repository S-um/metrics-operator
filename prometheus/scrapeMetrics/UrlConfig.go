package scrapeMetrics

import (
	"io"
	"net/http"
)

type UrlConfig struct {
	Protocol string `json:"protocol,omitempty"`
	Name     string `json:"name"`
	Path     string `json:"path,omitempty"`
	Port     string `json:"port,omitempty"`
}

func UrlArrToTargetConfigArr(s []UrlConfig) []TargetConfig {
	length := len(s)
	arr := make([]TargetConfig, length)
	for index := 0; index < length; index++ {
		arr[index] = &s[index]
	}
	return arr
}

func (u *UrlConfig) scrapeMetrics() (string, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	str, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func (u *UrlConfig) GetName() string {
	return u.Name
}

func (u *UrlConfig) GetPath() string {
	return u.Path
}

func (u *UrlConfig) GetPort() string {
	return u.Port
}

func (u *UrlConfig) GetProtocol() string {
	return u.Protocol
}

func (u *UrlConfig) String() string {
	return u.Protocol + "://" + u.Name + ":" + u.Port + "/" + u.Path
}
