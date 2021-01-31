package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post ...
type Post struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Type      Type
	Lead      int
	Lag       int
	Task      primitive.ObjectID
	Parent    primitive.ObjectID `bson:"parent, omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
