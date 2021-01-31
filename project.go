package model

import (
	"time"
)

// Project ...
type Project struct {
	ID           string
	Organisation Organisation
	Tag          string
	Owners       []User
	Name         string
	Tasks        []Task
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
