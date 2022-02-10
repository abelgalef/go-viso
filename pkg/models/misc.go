package models

type Constraints struct {
	Field         string `json:"-"`
	OperatorValue string `json:"-"`
	Value         string `json:"-"`
	Sort          string `json:"-"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
}
