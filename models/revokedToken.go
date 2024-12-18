package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RevokedToken struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Token string             `bson:"token"`
}
