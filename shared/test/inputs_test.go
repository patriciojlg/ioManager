package test

import (
	controllers "ioManager/controllers"
	"testing"
)

func TestInput(t *testing.T) {
	dummyInputFileArgs := getDummySaveInputFileArgsMap()

	batchInputObjectKey := controllers.SaveInput(dummyInputFileArgs)
	if batchInputObjectKey.StatusCode != 200 {

		t.Errorf("El path no es un S3 objectKey válido: %s", batchInputObjectKey.DetailError)
	}
	t.Logf("El id es: %s", batchInputObjectKey.DetailError)
}

func TestExplodeInput(t *testing.T) {
	dummyInputFileArgs := getDummySaveInputFileArgsMap()

	explodedInputObjectKey := controllers.ExplodeInput(dummyInputFileArgs)
	if explodedInputObjectKey.StatusCode != 200 {

		t.Errorf("El path no es un S3 objectKey válido: %s", explodedInputObjectKey.DetailError)
	}
	t.Logf("El id es: %s", explodedInputObjectKey.DetailError)
}
