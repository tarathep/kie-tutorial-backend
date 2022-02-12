package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tutorial is model data
type Tutorial struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Published   bool               `json:"published"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}
