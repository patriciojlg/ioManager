package models

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

type ConfigurationTemplate struct {
	SfnArn               string               `json:"sfn_arn"`
	TableBatch           string               `json:"table-batch"`
	TaskName             string               `json:"task-name"`
	FilePayload          FilePayload          `json:"file-payload"`
	ConcurrentExecutions *int                 `json:"concurrent-executions,omitempty"`
	MaxInflight          *int                 `json:"max-inflight,omitempty"`
	MaxRetries           *int                 `json:"max-retries,omitempty"`
	RetryInterval        *int                 `json:"retry-interval,omitempty"`
	AditionalValidation  *AditionalValidation `json:"aditional-validation,omitempty"`
}
type FilePayload struct {
	Columns []Column `json:"columns"`
}
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type AditionalValidation struct {
	Type      string `json:"type"`
	Arn       string `json:"arn"`
	EventArgs string `json:"event-args,omitempty"`
}

func (c *ConfigurationTemplate) FromJson(jsonOnDict map[string]any) error {
	if err := mapstructure.Decode(jsonOnDict, c); err != nil {
		return err
	}
	if c.TaskName == "" {
		return errors.New("task_name is required")
	}
	if c.TableBatch == "" {
		return errors.New("table-batch is required")
	}
	if c.SfnArn == "" {
		return errors.New("sfn_arn is required")
	}
	if c.FilePayload.Columns == nil {
		return errors.New("file-payload.columns is required")
	}
	if len(c.FilePayload.Columns) == 0 {
		return errors.New("file-payload.columns is empty")
	}
	return nil
}
