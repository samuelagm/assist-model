package model

import (
	"time"
)

// ProjectRequest holds a task operation
type ProjectRequest struct {
	Project struct {
		Tag            string   `yaml:"tag"`
		Name           string   `yaml:"name"`
		Start          string   `yaml:"start"`
		End            string   `yaml:"end"`
		Owners         []string `yaml:"owners"`
		UpdateInterval int      `yaml:"update-interval"`
	} `yaml:"project"`
	Tasks []struct {
		Name         string   `yaml:"name"`
		Owner        string   `yaml:"owner"`
		Milestones   []string `yaml:"milestones,omitempty"`
		Duration     int      `yaml:"duration,omitempty"`
		Deliverables []string `yaml:"deliverables,omitempty"`
		Waitfor      []string `yaml:"waitfor,omitempty"`
		Deliverable  []string `yaml:"deliverable,omitempty"`
	} `yaml:"tasks"`
}

// OperationRequest ...
type OperationRequest struct {
	Operation struct {
		Tag            string   `yaml:"tag"`
		Owners         []string `yaml:"owner"`
		UpdateInterval int      `yaml:"update-interval"`
		Triggers       []string `yaml:"triggers"`
	} `yaml:"operation"`
	Tasks []struct {
		Name         string   `yaml:"name"`
		Owner        string   `yaml:"owners"`
		Milestones   []string `yaml:"milestones,omitempty"`
		Duration     int      `yaml:"duration,omitempty"`
		Deliverables []string `yaml:"deliverables,omitempty"`
		Waitfor      []string `yaml:"waitfor,omitempty"`
		Deliverable  []string `yaml:"deliverable,omitempty"`
	} `yaml:"tasks"`
}

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
