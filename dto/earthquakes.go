package dto

type EarthquakeResponseDto struct {
	EarthquakesDto []EarthquakeDto `json:"earthquakes"`
}

type EarthquakeDto struct {
	EarthquakeID   string  `json:"earthquakeId"`
	PlaceName      string  `json:"placeName"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Depth          string  `json:"depth"`
	Magnitude      string  `json:"magnitude"`
	DateTime       string  `json:"dateTime"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
	AnnotationSize float32 `json:"annotationSize"`
}
