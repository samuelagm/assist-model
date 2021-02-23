package model

import (
	"time"
)

// Project ...
type Project struct {
	ID           string
	Organisation string
	Tag          string
	Owners       []string
	Name         string
	Tasks        []Task
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
