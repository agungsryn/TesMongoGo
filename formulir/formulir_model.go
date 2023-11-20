package formulir

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Formulir struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId"  bson:"user_id,omitempty"`
	Description string             `json:"description"  bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	ValidThru   time.Time          `json:"validThru" bson:"valid_thru"`
}

type FormulirRequest struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId"  bson:"user_id,omitempty"`
	Description string             `json:"description"  bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	ValidThru   string          `json:"validThru" bson:"valid_thru"`
}
