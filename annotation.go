package models

import (
	"time"
)

type LabelType string
type LabelValue string
type Classes map[string]float64
type RegressionRange [2]float64

const (
	ClassificationLabel LabelType = "classification"
	RegressionLabel     LabelType = "regression"
)

type Annotation struct {
	ID         string          `json:"id"`
	ImageID    string          `json:"image_id"` // ID of the associated image
	UserID     string          `json:"user_id"`  // ID of the user who created the annotation
	Polygons   string          `json:"polygons"` // Serialized polygon data (e.g., GeoJSON)
	LabelType  LabelType       `json:"label_type,omitempty"`
	LabelValue LabelValue      `json:"label_value,omitempty"`
	Classes    Classes         `json:"classes,omitempty"`
	Range      RegressionRange `json:"range,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ann *Annotation) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":          ann.ID,
		"image_id":    ann.ImageID,
		"user_id":     ann.UserID,
		"polygons":    ann.Polygons,
		"label_type":  ann.LabelType,
		"label_value": ann.LabelValue,
		"classes":     ann.Classes,
		"range":       ann.Range,
		"created_at":  ann.CreatedAt,
		"updated_at":  ann.UpdatedAt,
	}
}
