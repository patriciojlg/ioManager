package test

import (
	controllers "ioManager/controllers"
	"testing"
)

func TestImplodeOutput(t *testing.T) {
	dummyInputFileArgs := getDummyImploteOutputsArgs()

	batchInputObjectKey, err := controllers.ImplodeOutput(dummyInputFileArgs)
	if err != nil {

		t.Errorf("El path no es un S3 objectKey válido: %s", batchInputObjectKey.DetailError)
	}
	t.Logf("El id es: %s", batchInputObjectKey.DetailError)
}
