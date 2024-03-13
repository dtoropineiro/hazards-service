package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type FeatureCollectionResponse struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Id        int    `json:"id"`
	Comuna    string `json:"comuna"`
	Clave     string `json:"clave"`
	Ubicacion string `json:"ubicacion"`
	Fecha     string `json:"fecha"`
}

type Central123Client interface {
	GetLastFires() (*FeatureCollectionResponse, error)
}

func GetLastFires() (*FeatureCollectionResponse, error) {
	endpoint := viper.GetString("rest-services.central123.endpoint")

	resp, err := http.Get(fmt.Sprintf("%s/query?format=geojson", endpoint))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al obtener los últimos incendios: %s", resp.Status)
	}

	var fires FeatureCollectionResponse
	err = json.NewDecoder(resp.Body).Decode(&fires)
	if err != nil {
		return nil, err
	}

	return &fires, nil
}
