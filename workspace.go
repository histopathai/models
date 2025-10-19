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

// ToMap converts Workspace to map for Firestore
func (w *Workspace) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              w.ID,
		"name":            w.Name,
		"organ_type":      w.OrganType,
		"creator_id":      w.CreatorID,
		"description":     w.Description,
		"license":         w.License,
		"organization":    w.Organization,
		"resource_url":    w.ResourceURL,
		"release_year":    w.ReleaseYear,
		"release_version": w.ReleaseVersion,
		"created_at":      w.CreatedAt,
		"updated_at":      w.UpdatedAt,
	}
}
