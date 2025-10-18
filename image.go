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

type ImageInfo struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Format    string `json:"format"`
	SizeBytes int64  `json:"size_bytes"`
}

type Image struct {
	ID               string    `json:"id"`
	Filename         string    `json:"filename"`
	Info             ImageInfo `json:"info"`
	WorkspaceID      string    `json:"workspace_id"` // ID of the associated workspace
	GCSPath          string    `json:"gcs_path"`
	DZIGCSPath       string    `json:"dzi_gcs_path"`
	DZITilesGCSPath  string    `json:"dzi_tiles_gcs_path"`
	ThumbnailGCSPath string    `json:"thumbnail_gcs_path"`

	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Status    ImageStatus `json:"status"`
}

func (info *ImageInfo) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"width":      info.Width,
		"height":     info.Height,
		"format":     info.Format,
		"size_bytes": info.SizeBytes,
	}
}

func (img *Image) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":                 img.ID,
		"filename":           img.Filename,
		"info":               img.Info.ToJSON(),
		"workspace_id":       img.WorkspaceID,
		"gcs_path":           img.GCSPath,
		"dzi_gcs_path":       img.DZIGCSPath,
		"dzi_tiles_gcs_path": img.DZITilesGCSPath,
		"thumbnail_gcs_path": img.ThumbnailGCSPath,
		"created_at":         img.CreatedAt,
		"updated_at":         img.UpdatedAt,
		"status":             img.Status,
	}
}
