package model

import (
	"time"
)

// Task ...
type Task struct {
	ID           string
	Organisation string
	Name         string
	Owner        string
	Status       int
	Deliverables []string
	Duration     int
	Delta        int
	Milestones   []Milestone
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
