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

// FromMap populates Annotation from map data
func (a *Annotation) FromMap(data map[string]interface{}) {
	if v, ok := data["id"].(string); ok {
		a.ID = v
	}
	if v, ok := data["image_id"].(string); ok {
		a.ImageID = v
	}
	if v, ok := data["creator_id"].(string); ok {
		a.CreatorID = v
	}
	if v, ok := data["type_id"].(string); ok {
		a.TypeID = v
	}
	if v, ok := data["polygon"].([]interface{}); ok {
		points := make([]Point, len(v))
		for i, p := range v {
			if pointMap, ok := p.(map[string]interface{}); ok {
				var point Point
				if x, ok := pointMap["x"].(float64); ok {
					point.X = x
				}
				if y, ok := pointMap["y"].(float64); ok {
					point.Y = y
				}
				points[i] = point
			}
		}
		a.Polygon = points
	}
	if v, ok := data["class"].(string); ok {
		a.Class = &v
	}
	if v, ok := data["score"].(float64); ok {
		a.Score = &v
	}
	if v, ok := data["created_at"].(time.Time); ok {
		a.CreatedAt = v
	}
	if v, ok := data["updated_at"].(time.Time); ok {
		a.UpdatedAt = v
	}
}

func (at *AnnotationType) FromMap(data map[string]interface{}) {
	if v, ok := data["id"].(string); ok {
		at.ID = v
	}
	if v, ok := data["name"].(string); ok {
		at.Name = v
	}
	if v, ok := data["desc"].(string); ok {
		at.Desc = &v
	}
	if v, ok := data["score_enable"].(bool); ok {
		at.ScoreEnable = &v
	}
	if v, ok := data["score_name"].(string); ok {
		at.ScoreName = &v
	}
	if v, ok := data["score_range"].([]interface{}); ok && len(v) == 2 {
		var rangeArr [2]float64
		if val1, ok := v[0].(float64); ok {
			rangeArr[0] = val1
		}
		if val2, ok := v[1].(float64); ok {
			rangeArr[1] = val2
		}
		at.ScoreRange = rangeArr
	}
	if v, ok := data["classification_enable"].(bool); ok {
		at.ClassificationEnable = &v
	}
	if v, ok := data["class_list"].([]interface{}); ok {
		classList := make([]string, len(v))
		for i, class := range v {
			if classStr, ok := class.(string); ok {
				classList[i] = classStr
			}
		}
		at.ClassList = classList
	}
	if v, ok := data["created_at"].(time.Time); ok {
		at.CreatedAt = v
	}
	if v, ok := data["updated_at"].(time.Time); ok {
		at.UpdatedAt = v
	}
}
