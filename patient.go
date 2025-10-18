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
