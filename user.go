package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User ...
type User struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Organisation primitive.ObjectID
	UserName     string
	FirstName    string
	LastName     string
	Email        string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
