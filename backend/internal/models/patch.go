package models

type Patch struct {
	ID            string        `json:"id"`
	Name          string        `json:"title"`
	ViewportState ViewportState `json:"viewportState"`
	Modules       []Module      `json:"modules"`
	Connections   []Connection  `json:"connections"`
}

type ViewportState struct {
	Scale   float64 `json:"scale"`
	OffsetX float64 `json:"offsetX"`
	OffsetY float64 `json:"offsetY"`
}

type Module struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	PositionX  float64                `json:"positionX"`
	PositionY  float64                `json:"positionY"`
	Parameters map[string]interface{} `json:"parameters"`
}

type Connection struct {
	ID           string `json:"id"`
	SourceModule string `json:"sourceModule"`
	SourceHandle string `json:"sourceHandle"`
	TargetModule string `json:"targetModule"`
	TargetHandle string `json:"targetHandle"`
}
