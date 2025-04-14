package models

type MetaDataInput struct {
	TaskID     string `json:"task_id"`
	S3Key      string `json:"s3_key"`
	Format     string `json:"format"`
	UploadedAt int64  `json:"uploaded_at"`
}
