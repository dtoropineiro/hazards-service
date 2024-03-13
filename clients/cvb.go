package clients

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type EmergencyResponse struct {
	Bo   bool        `json:"bo"`
	Data []Emergency `json:"data"`
}

type Emergency struct {
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
	Dir    string `json:"dir"`
	Sigla  string `json:"sigla"`
	Nombre string `json:"nombre"`
	Fecha  string `json:"fecha"`
}

type CvbClient interface {
	GetLastFireEmergencies() (*EmergencyResponse, error)
}

func GetLastFireEmergencies() (*EmergencyResponse, error) {
	cbvEndpoint := viper.GetString("rest-services.cbv.endpoint")
	response, err := http.Get(fmt.Sprintf("%s?type=13", cbvEndpoint))
	if err != nil {
		logrus.Error(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error(err)
	}

	var emergencyResponse EmergencyResponse
	err = json.Unmarshal(body, &emergencyResponse)
	if err != nil {
		logrus.Error(err)
	}

	return &emergencyResponse, nil
}
