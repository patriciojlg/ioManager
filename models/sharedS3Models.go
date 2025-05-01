package models

type S3DownloadObjectKeyResult struct {
	Key  string
	Body []byte
	Err  error
}
