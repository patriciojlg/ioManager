package models

type Event struct {
	Command string         `json:"command"`
	Args    map[string]any `json:"args"`
}
