package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"ioManager/models"
	settings "ioManager/settings"
	"ioManager/shared/utils"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3

func init() {
	s3Client = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(settings.AWS_REGION),
	})))
}
func listObjectsS3(prefix string) ([]string, error) {
	resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(settings.BUCKETNAME),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("error listing objects in S3: %w", err)
	}

	var keys []string
	for _, obj := range resp.Contents {
		keys = append(keys, *obj.Key)
	}

	return keys, nil
}
func uploadJsonToS3(taskID, filename string, data any) (string, error) {
	bucket := settings.BUCKETNAME
	key := fmt.Sprintf("%s/%s", taskID, filename)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %w", err)
	}

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return "", fmt.Errorf("error uploading to S3: %w", err)
	}

	return key, nil
}
func uploadBytesToS3(key string, data []byte) (string, error) {
	bucket := settings.BUCKETNAME
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return "", fmt.Errorf("error uploading to S3: %w", err)
	}

	return key, nil
}
func getObjectS3(key string) ([]byte, error) {
	bucket := settings.BUCKETNAME
	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting object from S3: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}
func SaveInputFileOnS3(key string, data []byte) error {
	_, err := uploadBytesToS3(key, data)
	if err != nil {
		return fmt.Errorf("error uploading input file to S3: %w", err)
	}
	return nil
}
func SaveMetadataOnS3(taskID string, metadata any) error {
	_, err := uploadJsonToS3(taskID, "metadata.json", metadata)
	if err != nil {
		return fmt.Errorf("error uploading metadata to S3: %w", err)
	}
	return nil
}

func GetMainInputFileFromS3(args models.SaveInputFileArgs) ([]byte, error) {
	prefixMainInput := utils.GetPrefixMainInputS3Object(args)
	files, err := listObjectsS3(prefixMainInput)
	if err != nil {
		return nil, fmt.Errorf("error listing objects in S3: %w", err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no files found in S3 for taskID: %s", args.Id)
	}
	// search xlsx file
	var key string
	for _, file := range files {
		if file[len(file)-5:] == ".xlsx" {
			key = file
			break
		}
	}

	data, err := getObjectS3(key)
	if err != nil {
		return nil, fmt.Errorf("error getting input file from S3: %w", err)
	}
	return data, nil
}
func UploadJSONFilesConcurrentlyV1(files []models.JSONFile, prefix string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(files))
	semaphore := make(chan struct{}, settings.S3_MAX_WORKERS)

	for _, file := range files {
		wg.Add(1)

		go func(f models.JSONFile) {
			defer wg.Done()
			semaphore <- struct{}{} // ocupar slot

			key := fmt.Sprintf("%s/%s", prefix, f.Filename)
			_, err := s3Client.PutObject(&s3.PutObjectInput{
				Bucket:      aws.String(settings.BUCKETNAME),
				Key:         aws.String(key),
				Body:        bytes.NewReader(f.Data),
				ContentType: aws.String("application/json"),
			})
			if err != nil {
				errCh <- fmt.Errorf("error subiendo %s: %w", f.Filename, err)
			} else {
				log.Printf("✅ Subido: %s", key)
			}

			<-semaphore // liberar slot
		}(file)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		for err := range errCh {
			log.Println("❌", err)
		}
		return fmt.Errorf("algunos archivos fallaron en la subida")
	}

	return nil
}
