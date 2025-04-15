package controllers

import (
	"ioManager/models"
	providers "ioManager/providers"
	"ioManager/shared/utils"
)

func SaveInput(args map[string]any) models.Response {
	argsSaveInput, err := models.SaveInputFileArgsFromArgs(args)
	if err != nil {
		return models.Error400Response(err)
	}
	decoded, err := utils.AnyInputBodyEncodedToBytes(argsSaveInput.Format, argsSaveInput.Body)
	if err != nil {
		return models.Error400Response(err)
	}
	// Subir a S3
	key, err := utils.GetMainInputFileBatchS3ObjectKey(argsSaveInput)
	if err != nil {
		return models.Error400Response(err)
	}
	err = providers.SaveInputFileOnS3(key, decoded)
	if err != nil {
		return models.Error400Response(err)
	}
	meta := models.MetaDataInput{
		TaskID:     argsSaveInput.Id,
		S3Key:      key,
		Format:     argsSaveInput.Format,
		UploadedAt: 0, // TODO: Cambiar por timestamp
	}

	return models.Response{
		StatusCode:  200,
		Body:        meta,
		Error:       false,
		DetailError: "",
	}

}

/*
func ExplodeInput(args map[string]any) models.Response {
	data, err := providers.GetMainInputFileFromS3(args)
	if err != nil {
		return models.Error400Response(err)
	}
	rows, error := providers.ReadExcelFromBytes(data)
	if error != nil {
		return models.Error400Response(error)
	}
	if len(rows) < 2 {
		return models.Error400Response(fmt.Errorf("no hay suficientes filas en el archivo"))
	}
	explodedRows, err := providers.RowsToArrayJsonStruct(rows)
	if err != nil {
		return models.Error400Response(err)
	}
	prefixExplodedS3 := utils.GetMainInputExplodedBatchPrefixS3ObjectKey(args["id"].(string), args["task_name"].(string))
	err = providers.UploadJSONFilesConcurrentlyV1(explodedRows, prefixExplodedS3)
	if err != nil {
		return models.Error400Response(err)
	}
	return models.Response{
		StatusCode:  200,
		Body:        nil,
		Error:       false,
		DetailError: "",
	}
}
*/
