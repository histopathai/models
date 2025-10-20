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

var SupportedImageFormats = []string{"jpeg", "png", "tiff", "bmp", "gif", "svs", "ndpi", "dng", "png", "jpg", "tif"}

type Image struct {
	ID            string      `firestore:"id"`
	FileName      string      `firestore:"filename"`
	Width         int         `firestore:"width"`
	Height        int         `firestore:"height"`
	Format        string      `firestore:"format"`
	SizeBytes     int64       `firestore:"size_bytes"`
	CreatorID     string      `firestore:"creator_id"`   // ID of the user who uploaded the image
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
		"file_name":      img.FileName,
		"width":          img.Width,
		"height":         img.Height,
		"format":         img.Format,
		"size_bytes":     img.SizeBytes,
		"creator_id":     img.CreatorID,
		"workspace_id":   img.WorkspaceID,
		"patient_id":     img.PatientID,
		"origin_path":    img.OriginPath,
		"processed_path": img.ProcessedPath,
		"created_at":     img.CreatedAt,
		"updated_at":     img.UpdatedAt,
		"status":         img.Status,
	}
}

// From Map
func (img *Image) FromMap(data map[string]interface{}) {
	if v, ok := data["id"].(string); ok {
		img.ID = v
	}
	if v, ok := data["file_name"].(string); ok {
		img.FileName = v
	}
	if v, ok := data["width"].(int64); ok {
		img.Width = int(v)
	}
	if v, ok := data["height"].(int64); ok {
		img.Height = int(v)
	}
	if v, ok := data["format"].(string); ok {
		img.Format = v
	}
	if v, ok := data["size_bytes"].(int64); ok {
		img.SizeBytes = v
	}
	if v, ok := data["creator_id"].(string); ok {
		img.CreatorID = v
	}
	if v, ok := data["workspace_id"].(string); ok {
		img.WorkspaceID = v
	}
	if v, ok := data["patient_id"].(string); ok {
		img.PatientID = v
	}
	if v, ok := data["origin_path"].(string); ok {
		img.OriginPath = v
	}
	if v, ok := data["processed_path"].(string); ok {
		img.ProcessedPath = v
	}
	if v, ok := data["created_at"].(time.Time); ok {
		img.CreatedAt = v
	}
	if v, ok := data["updated_at"].(time.Time); ok {
		img.UpdatedAt = v
	}
	if v, ok := data["status"].(string); ok {
		img.Status = ImageStatus(v)
	}
}
