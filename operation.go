package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Operation ...
type Operation struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Organisation primitive.ObjectID
	Tag          string
	Trigger      string
	Owners       []primitive.ObjectID
	Name         string
	Tasks        []primitive.ObjectID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
