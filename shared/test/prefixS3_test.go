package test

import (
	models "ioManager/models"
	utils "ioManager/shared/utils"
	"regexp"
	"strings"
	"testing"
)

func isS3Path(s3ObjectKey string) bool {
	if strings.Contains(s3ObjectKey, "//") {
		return false
	}
	regex := regexp.MustCompile(`[a-z0-9\-\.]{3,63}/.+`)
	return regex.MatchString(s3ObjectKey)
}
func TestBatchInputS3ObjectKey(t *testing.T) {
	x, err := models.SaveInputFileArgsFromArgs(map[string]any{
		"task_name":    "solicitud-transferencia",
		"account_name": "prenda-chile",
		"id":           "r52d434",
		"body":         "zxczxczxc",
		"filename":     "file_predachile_01.xlsx",
		"format":       "xlsx",
		"coded":        "base64",
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	batchInputObjectKey, err := utils.GetMainInputFileBatchS3ObjectKey(x)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !isS3Path(batchInputObjectKey) {
		t.Errorf("El path no es un S3 objectKey válido: %s", batchInputObjectKey)
	}
	t.Logf("El id es: %s", batchInputObjectKey)
}
