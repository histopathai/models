// models/patient.go (Yeni dosya)
package models

import "time"

type Patient struct {
	ID        string    `firestore:"id,omitempty"` // Firestore ID veya kendi ID'niz
	Age       *int      `firestore:"age,omitempty"`
	Gender    *string   `firestore:"gender,omitempty"`
	Race      *string   `firestore:"race,omitempty"`
	Disease   *string   `firestore:"disease,omitempty"`
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
		"history":    p.History,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
	}
}
