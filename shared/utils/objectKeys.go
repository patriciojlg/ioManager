package utils

import "errors"

func GetMainInputBatchS3ObjectKey(taskID, taskName, format string) (string, error) {
	// example: "batch-payloads/batch-input/{task_name}/{name_account}/{uuid}/{uuid}.json"
	switch format {
	case "json":
		return "batch-payloads/batch-input/" + taskName + "/" + taskID + "/" + taskID + ".json", nil
	case "csv":
		return "batch-payloads/batch-input/" + taskName + "/" + taskID + "/" + taskID + ".csv", nil
	case "xlsx":
		return "batch-payloads/batch-input/" + taskName + "/" + taskID + "/" + taskID + ".xlsx", nil
	default:
		return "", errors.New("formato no soportado: " + format)
	}

}
func GetMainInputBatchPrefixS3ObjectKey(taskID, taskName string) string {
	// example: "batch-payloads/batch-input/{task_name}/{name_account}/{uuid}/"
	return "batch-payloads/batch-input/" + taskName + "/" + taskID + "/"
}

func GetMainInputExplodedBatchPrefixS3ObjectKey(taskID, taskName string) string {
	// example: "batch-payloads/batch-input/{task_name}/{name_account}/{uuid}/"
	return "batch-payloads/batch-input/" + taskName + "/" + taskID + "/exploded/"
}
