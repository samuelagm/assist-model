package model

import (
	"time"
)

// Dependency ...
type Dependency struct {
	Type      Type
	Lead      int
	Lag       int
	Task      Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Milestone ...
type Milestone struct {
	Name   string
	Status Status
}
