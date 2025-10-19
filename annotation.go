package models

import (
	"time"
)

type Point struct {
	X float64 `firestore:"x"`
	Y float64 `firestore:"y"`
}

type AnnotationType struct {
	ID                   string     `firestore:"id,omitempty"`
	Name                 string     `firestore:"name"`
	Desc                 *string    `firestore:"desc"`
	ScoreEnable          *bool      `firestore:"score_enable"`
	ScoreName            *string    `firestore:"score_name"`
	ScoreRange           [2]float64 `firestore:"score_range"`
	ClassificationEnable *bool      `firestore:"classification_enable"`
	ClassList            []string   `firestore:"class_list"`
	CreatedAt            time.Time  `firestore:"created_at"`
	UpdatedAt            time.Time  `firestore:"updated_at"`
}

type Annotation struct {
	ID        string    `firestore:"id,omitempty"`
	ImageID   string    `firestore:"image_id"`        // ID of the associated image
	CreatorID string    `firestore:"creator_id"`      // ID of the user who created the annotation
	TypeID    string    `firestore:"type_id"`         // ID of the annotation type
	Polygon   []Point   `firestore:"polygon"`         // Serialized polygon data (e.g., GeoJSON)
	Class     *string   `firestore:"class,omitempty"` // Class label for the annotation
	Score     *float64  `firestore:"score,omitempty"` // Score value for the annotation
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

func (p *Point) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"x": p.X,
		"y": p.Y,
	}
}

// ToMap converts Annotation to map for Firestore
func (a *Annotation) ToMap() map[string]interface{} {
	polygonMaps := make([]map[string]interface{}, len(a.Polygon))
	for i, p := range a.Polygon {
		polygonMaps[i] = p.ToMap()
	}

	return map[string]interface{}{
		"id":         a.ID,
		"image_id":   a.ImageID,
		"creator_id": a.CreatorID,
		"type_id":    a.TypeID,
		"polygon":    polygonMaps,
		"class":      a.Class,
		"score":      a.Score,
		"created_at": a.CreatedAt,
		"updated_at": a.UpdatedAt,
	}
}

func (at *AnnotationType) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                    at.ID,
		"name":                  at.Name,
		"desc":                  at.Desc,
		"score_enable":          at.ScoreEnable,
		"score_name":            at.ScoreName,
		"score_range":           at.ScoreRange,
		"classification_enable": at.ClassificationEnable,
		"class_list":            at.ClassList,
		"created_at":            at.CreatedAt,
		"updated_at":            at.UpdatedAt,
	}
}
