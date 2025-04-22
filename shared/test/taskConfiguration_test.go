package test

import (
	"fmt"
	"ioManager/controllers"
	"testing"
)

func TestGetTaskConfiguration(t *testing.T) {
	dummyInputFileArgs := getDummyGetTaskConfigurationArgs()

	taskConfiguration, err := controllers.GetConfigurationTemplate(dummyInputFileArgs.TaskName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Println(taskConfiguration)
}
