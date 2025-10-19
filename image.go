package models

import (
	"time"
)

type ImageStatus string

const (
	StatusUploadWaiting    ImageStatus = "upload_waiting"
	StatusUploaded         ImageStatus = "uploaded"
	StatusUploadFailed     ImageStatus = "upload_failed"
	StatusProcessing       ImageStatus = "processing"
	StatusProcessed        ImageStatus = "processed"
	StatusProcessingFailed ImageStatus = "processing_failed"
)

type Image struct {
	ID            string      `firestore:"id"`
	Filename      string      `firestore:"filename"`
	Width         int         `firestore:"width"`
	Height        int         `firestore:"height"`
	Format        string      `firestore:"format"`
	SizeBytes     int64       `firestore:"size_bytes"`
	WorkspaceID   string      `firestore:"workspace_id"` // ID of the associated workspace
	PatientID     string      `firestore:"patient_id,omitempty"`
	OriginPath    string      `firestore:"origin_path,omitempty"`
	ProcessedPath string      `firestore:"processed_path,omitempty"`
	CreatedAt     time.Time   `firestore:"created_at"`
	UpdatedAt     time.Time   `firestore:"updated_at"`
	Status        ImageStatus `firestore:"status"`
}

// ToMap converts Image to map for Firestore
func (img *Image) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             img.ID,
		"filename":       img.Filename,
		"width":          img.Width,
		"height":         img.Height,
		"format":         img.Format,
		"size_bytes":     img.SizeBytes,
		"workspace_id":   img.WorkspaceID,
		"patient_id":     img.PatientID,
		"origin_path":    img.OriginPath,
		"processed_path": img.ProcessedPath,
		"created_at":     img.CreatedAt,
		"updated_at":     img.UpdatedAt,
		"status":         img.Status,
	}
}
