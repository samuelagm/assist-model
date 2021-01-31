package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organisation ...
type Organisation struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
