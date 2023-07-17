package models

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}
