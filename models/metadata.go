package models

type MetaDataInput struct {
	TaskID     string `json:"task_id"`
	S3Key      string `json:"s3_key"`
	Format     string `json:"format"`
	UploadedAt int64  `json:"uploaded_at"`
}

type MetaDataExplodedInput struct {
	TaskID         string `json:"task_id"`
	TaskName       string `json:"task_name"`
	PrefixExploded string `json:"prefix_exploded"`
	Format         string `json:"format"`
	QuantityFiles  int    `json:"quantity_files"`
	UploadedAt     int64  `json:"uploaded_at"`
}
