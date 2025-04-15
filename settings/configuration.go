package settings

const (
	// AWS SETTINGS
	BUCKETNAME = "batchers-sl"
	AWS_REGION = "us-east-1"

	// GO ROUTINES
	S3_MAX_WORKERS = 10
	// S3 - PREFIX  AND PARTITIONS

	BASE_BATCH_PREFIX_S3   = "batch-payloads/"
	MAIN_INPUT_FOLDER      = "batch-input/"
	EXPLODED_INPUT_FOLDER  = "exploded-input/"
	EXPLODED_OUTPUT_FOLDER = "exploded-output/"
	IMPLODED_OUTPUT_FOLDER = "imploded-output/"
)
