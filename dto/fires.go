package dto

type FireResponseDto struct {
	FireDto []FireDto `json:"fires"`
}

type FireDto struct {
	Id             int     `json:"id"`
	FireId         string  `json:"fireId"`
	PlaceName      string  `json:"placeName"`
	District       string  `json:"district"`
	Key            string  `json:"key"`
	KeyDescription string  `json:"keyDescription"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	DateTime       string  `json:"dateTime"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
}
