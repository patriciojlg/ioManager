package utils

import (
	"errors"
	"ioManager/models"
	"ioManager/settings"
)

func getBaseTaskNameAccountS3Object(argsTask models.SaveInputFileArgs) string {
	return settings.BASE_BATCH_PREFIX_S3 + argsTask.AccountName + "/" + argsTask.TaskName + "/" + argsTask.Id + "/"
}

func GetMainInputFileBatchS3ObjectKey(argsTask models.SaveInputFileArgs) (string, error) {
	baseS3ObjectKey := GetPrefixMainInputS3Object(argsTask)
	mainInputFileLessFormat := baseS3ObjectKey + argsTask.Id
	switch argsTask.Format {
	case "json":
		return mainInputFileLessFormat + ".json", nil
	case "csv":
		return mainInputFileLessFormat + ".csv", nil
	case "xlsx":
		return mainInputFileLessFormat + ".xlsx", nil
	default:
		return "", errors.New("formato no soportado: " + argsTask.Format)
	}

}
func GetMainFileImplodeOutput(argsTask models.SaveInputFileArgs) (string, error) {
	baseS3ObjectKey := GetPrefixImplodedOutput(argsTask)
	mainInputFileLessFormat := baseS3ObjectKey + argsTask.Id
	switch argsTask.Format {
	case "json":
		return mainInputFileLessFormat + ".json", nil
	case "csv":
		return mainInputFileLessFormat + ".csv", nil
	case "xlsx":
		return mainInputFileLessFormat + ".xlsx", nil
	default:
		return "", errors.New("formato no soportado: " + argsTask.Format)
	}
}
func GetPrefixMainInputS3Object(argsTask models.SaveInputFileArgs) string {
	baseTaskNameAccount := getBaseTaskNameAccountS3Object(argsTask)
	return baseTaskNameAccount + settings.MAIN_INPUT_FOLDER
}
func GetPrefixExplodedInputs(argsTask models.SaveInputFileArgs) string {
	baseTaskNameAccount := getBaseTaskNameAccountS3Object(argsTask)
	return baseTaskNameAccount + settings.EXPLODED_INPUT_FOLDER
}

func GetPrefixExplodedOutputs(argsTask models.SaveInputFileArgs) string {
	baseTaskNameAccount := getBaseTaskNameAccountS3Object(argsTask)
	return baseTaskNameAccount + settings.EXPLODED_OUTPUT_FOLDER
}

func GetPrefixImplodedOutput(argsTask models.SaveInputFileArgs) string {
	baseTaskNameAccount := getBaseTaskNameAccountS3Object(argsTask)
	return baseTaskNameAccount + settings.IMPLODED_OUTPUT_FOLDER
}
