package model

import (
	"time"
)

// Task ...
type Task struct {
	ID           string
	Organisation Organisation
	Name         string
	Owner        User
	Status       Status
	Dependencies []Dependency
	Deliverables []string
	Duration     int
	Delta        int
	Milestones   []Milestone
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
