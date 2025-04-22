package controllers

import (
	"errors"
	"ioManager/models"
	providers "ioManager/providers"
)

func isValidTaskName(taskname string) (bool, error) {
	task_names, err := providers.ListSettingsTaskFilesOnS3()
	if err != nil {
		return false, err
	}
	if taskname == "" {
		return false, nil
	}
	for _, task := range task_names {
		if task == taskname {
			return true, nil
		}
	}
	return false, nil

}
func GetValidTaskNames() (models.Response, error) {
	task_names, err := providers.ListSettingsTaskFilesOnS3()
	if err != nil {
		return models.Error400Response(err), err
	}
	if len(task_names) == 0 {
		return models.Error400Response(err), errors.New("no task names found")
	}
	return models.Response{
		StatusCode:  200,
		Body:        []string(task_names),
		Error:       false,
		DetailError: "",
	}, nil
}
func GetConfigurationTemplate(taskName string) (models.Response, error) {
	isValid, err := isValidTaskName(taskName)
	if err != nil {
		return models.Error400Response(err), err
	}
	if !isValid {
		return models.Error400Response(errors.New("task name not found")), errors.New("task name not found")
	}

	configurationTemplate, err := providers.GetJsonBatchSettings(taskName)
	if err != nil {
		return models.Error400Response(err), err
	}
	return models.Response{
		StatusCode:  200,
		Body:        configurationTemplate,
		Error:       false,
		DetailError: "",
	}, nil
}
