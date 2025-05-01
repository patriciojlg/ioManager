package models

import (
	"ioManager/settings"

	"github.com/mitchellh/mapstructure"
)

type ImplodeOutputFilesArgs struct {
	TaskName    string `mapstructure:"task_name"`
	AccountName string `mapstructure:"account_name"`
	Id          string `mapstructure:"id"`
	Format      string `mapstructure:"format"`
}

func ImplodeOutputFilesArgsFromArgs(args map[string]any) (ImplodeOutputFilesArgs, error) {
	c := &ImplodeOutputFilesArgs{}
	if err := mapstructure.Decode(args, c); err != nil {
		return *c, err
	}
	if c.TaskName == "" {
		return *c, settings.ErrTaskNameRequired
	}
	if c.Id == "" {
		return *c, settings.ErrIdRequired
	}
	if c.AccountName == "" {
		return *c, settings.ErrAccountNameRequired
	}
	if c.Format == "" {
		return *c, settings.ErrFormatRequired
	}
	return *c, nil
}
