package domain

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
