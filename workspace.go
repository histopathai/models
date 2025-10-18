package models

import (
	"time"
)

type Workspace struct {
	ID             string    `firestore:"id,omitempty"`
	Name           string    `firestore:"name"`
	OrganType      string    `firestore:"organ_type"`
	CreatorID      string    `firestore:"creator_id"` // User ID of the creator
	Description    string    `firestore:"description"`
	License        string    `firestore:"license"`
	Organization   string    `firestore:"organization"`
	ResourceURL    string    `firestore:"resource_url"`
	ReleaseYear    int       `firestore:"release_year,omitempty"`
	ReleaseVersion string    `firestore:"release_version,omitempty"`
	CreatedAt      time.Time `firestore:"created_at"`
	UpdatedAt      time.Time `firestore:"updated_at"`
}
