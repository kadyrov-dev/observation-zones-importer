package types

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Polygon    `json:"geometry"`
}

type Properties struct {
	ZoneID     int64  `json:"zone_id"`
	ZoneTypeID int    `json:"zone_type_id"`
	Address    string `json:"address"`
}

type Polygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}
