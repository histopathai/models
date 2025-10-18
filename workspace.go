package models

import (
	"time"
)

type Workspace struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	OrganType   string `json:"organ_type"`
	CreatorID   string `json:"creator_id"` // User ID of the creator
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ws *Workspace) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":          ws.ID,
		"name":        ws.Name,
		"organ_type":  ws.OrganType,
		"creator_id":  ws.CreatorID,
		"description": ws.Description,
		"created_at":  ws.CreatedAt,
		"updated_at":  ws.UpdatedAt,
	}
}
