package models

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

type SaveInputFileArgs struct {
	Body     string `mapstructure:"body"`
	Filename string `mapstructure:"filename"`
	Format   string `mapstructure:"format"`
	Coded    string `mapstructure:"coded"`
	Id       string `mapstructure:"id"`
	TaskName string `mapstructure:"task_name"`
}

// FromArgs llena el struct desde un map[string]any
func SaveInputFileArgsFromArgs(args map[string]any) (SaveInputFileArgs, error) {
	c := &SaveInputFileArgs{}
	if err := mapstructure.Decode(args, c); err != nil {
		return *c, err
	}
	if c.Coded == "" {
		return *c, errors.New("coded is required")
	}
	if c.TaskName == "" {
		return *c, errors.New("task_name is required")
	}
	// Validaciones simples (podés usar una librería como validator si querés)
	if c.Body == "" {
		return *c, errors.New("name is required")
	}
	if c.Filename == "" {
		return *c, errors.New("filename is required")
	}
	if c.Format == "" {
		return *c, errors.New("format is required")
	}
	if c.Id == "" {
		return *c, errors.New("id is required")
	}

	return *c, nil
}
