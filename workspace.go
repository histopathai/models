package models

import (
	"time"
)

type Workspace struct {
	ID               string    `firestore:"id,omitempty"`
	Name             string    `firestore:"name"`
	OrganType        string    `firestore:"organ_type"`
	CreatorID        string    `firestore:"creator_id"` // User ID of the creator
	Description      string    `firestore:"description"`
	License          string    `firestore:"license"`
	Organization     string    `firestore:"organization"`
	ResourceURL      string    `firestore:"resource_url"`
	ReleaseYear      int       `firestore:"release_year,omitempty"`
	ReleaseVersion   string    `firestore:"release_version,omitempty"`
	AnnotationTypeID string    `firestore:"annotation_type_id,omitempty"`
	CreatedAt        time.Time `firestore:"created_at"`
	UpdatedAt        time.Time `firestore:"updated_at"`
}

// ToMap converts Workspace to map for Firestore
func (w *Workspace) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                 w.ID,
		"name":               w.Name,
		"organ_type":         w.OrganType,
		"creator_id":         w.CreatorID,
		"description":        w.Description,
		"license":            w.License,
		"organization":       w.Organization,
		"resource_url":       w.ResourceURL,
		"release_year":       w.ReleaseYear,
		"release_version":    w.ReleaseVersion,
		"annotation_type_id": w.AnnotationTypeID,
		"created_at":         w.CreatedAt,
		"updated_at":         w.UpdatedAt,
	}
}

// FromMap populates Workspace from map data
func (w *Workspace) FromMap(data map[string]interface{}) {
	if v, ok := data["id"].(string); ok {
		w.ID = v
	}
	if v, ok := data["name"].(string); ok {
		w.Name = v
	}
	if v, ok := data["organ_type"].(string); ok {
		w.OrganType = v
	}
	if v, ok := data["creator_id"].(string); ok {
		w.CreatorID = v
	}
	if v, ok := data["description"].(string); ok {
		w.Description = v
	}
	if v, ok := data["license"].(string); ok {
		w.License = v
	}
	if v, ok := data["organization"].(string); ok {
		w.Organization = v
	}
	if v, ok := data["resource_url"].(string); ok {
		w.ResourceURL = v
	}
	if v, ok := data["annotation_type_id"].(string); ok {
		w.AnnotationTypeID = v
	}
	if v, ok := data["release_year"].(int64); ok {
		w.ReleaseYear = int(v)
	}
	if v, ok := data["release_version"].(string); ok {
		w.ReleaseVersion = v
	}
	if v, ok := data["created_at"].(time.Time); ok {
		w.CreatedAt = v
	}
	if v, ok := data["updated_at"].(time.Time); ok {
		w.UpdatedAt = v
	}
}
