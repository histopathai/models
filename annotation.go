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
