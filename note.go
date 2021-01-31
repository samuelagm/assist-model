package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Note ...
type Note struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Organisation primitive.ObjectID
	Task         primitive.ObjectID
	Owner        primitive.ObjectID
	message      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
