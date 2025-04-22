package models

type GetTaskConfigurationArgs struct {
	TaskName    string `mapstructure:"task_name"`
	AccountName string `mapstructure:"account_name"`
	Id          string `mapstructure:"id"`
}
