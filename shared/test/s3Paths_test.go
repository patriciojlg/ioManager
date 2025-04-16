package test

import (
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
	dummyInputFileArgs := getDummySaveInputFileArgs()

	batchInputObjectKey, err := utils.GetMainInputFileBatchS3ObjectKey(dummyInputFileArgs)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !isS3Path(batchInputObjectKey) {
		t.Errorf("El path no es un S3 objectKey válido: %s", batchInputObjectKey)
	}
	t.Logf("El id es: %s", batchInputObjectKey)
}

func TestExplodedInputS3ObjectKey(t *testing.T) {
	dummyInputFileArgs := getDummySaveInputFileArgs()

	explodedInputObjectKey := utils.GetPrefixExplodedInputs(dummyInputFileArgs)

	if !isS3Path(explodedInputObjectKey) {
		t.Errorf("El path no es un S3 objectKey válido: %s", explodedInputObjectKey)
	}
	t.Logf("El id es: %s", explodedInputObjectKey)
}

func TestPrefixExplodedOutputS3ObjectKey(t *testing.T) {
	dummyInputFileArgs := getDummySaveInputFileArgs()

	explodedOutputObjectKey := utils.GetPrefixExplodedOutputs(dummyInputFileArgs)

	if !isS3Path(explodedOutputObjectKey) {
		t.Errorf("El path no es un S3 objectKey válido: %s", explodedOutputObjectKey)
	}
	t.Logf("El id es: %s", explodedOutputObjectKey)
}

func TestPrefixImplodedOutputS3ObjectKey(t *testing.T) {
	dummyInputFileArgs := getDummySaveInputFileArgs()

	implodedOutputObjectKey := utils.GetPrefixImplodedOutput(dummyInputFileArgs)

	if !isS3Path(implodedOutputObjectKey) {
		t.Errorf("El path no es un S3 objectKey válido: %s", implodedOutputObjectKey)
	}
	t.Logf("El id es: %s", implodedOutputObjectKey)
}

func TestGetMainFileImplodeOutput(t *testing.T) {
	dummyInputFileArgs := getDummySaveInputFileArgs()

	implodedOutputObjectKey, err := utils.GetMainFileImplodeOutput(dummyInputFileArgs)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !isS3Path(implodedOutputObjectKey) {
		t.Errorf("El path no es un S3 objectKey válido: %s", implodedOutputObjectKey)
	}
	t.Logf("El id es: %s", implodedOutputObjectKey)
}
