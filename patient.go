// models/patient.go (Yeni dosya)
package models

import "time"

type Patient struct {
	ID        string    `firestore:"id,omitempty"` // Firestore ID veya kendi ID'niz
	Age       *int      `firestore:"age,omitempty"`
	Gender    *string   `firestore:"gender,omitempty"`
	Race      *string   `firestore:"race,omitempty"`
	Disease   *string   `firestore:"disease,omitempty"`
	SubType   *string   `firestore:"subtype,omitempty"`
	Grade     *string   `firestore:"grade,omitempty"`
	History   *string   `firestore:"history,omitempty"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

// ToMap converts Patient to map for Firestore
func (p *Patient) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         p.ID,
		"age":        p.Age,
		"gender":     p.Gender,
		"race":       p.Race,
		"disease":    p.Disease,
		"subtype":    p.SubType,
		"grade":      p.Grade,
		"history":    p.History,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
	}
}

// FromMap populates Patient from map data
func (p *Patient) FromMap(data map[string]interface{}) {
	if v, ok := data["id"].(string); ok {
		p.ID = v
	}
	if v, ok := data["age"].(int64); ok {
		age := int(v)
		p.Age = &age
	}
	if v, ok := data["gender"].(string); ok {
		p.Gender = &v
	}
	if v, ok := data["race"].(string); ok {
		p.Race = &v
	}
	if v, ok := data["disease"].(string); ok {
		p.Disease = &v
	}
	if v, ok := data["subtype"].(string); ok {
		p.SubType = &v
	}
	if v, ok := data["grade"].(string); ok {
		p.Grade = &v
	}
	if v, ok := data["history"].(string); ok {
		p.History = &v
	}
	if v, ok := data["created_at"].(time.Time); ok {
		p.CreatedAt = v
	}
	if v, ok := data["updated_at"].(time.Time); ok {
		p.UpdatedAt = v
	}
}
