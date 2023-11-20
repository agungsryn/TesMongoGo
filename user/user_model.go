package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID    `json:"id"  bson:"_id,omitempty"`
	Nama      string    `json:"nama"  bson:"nama,omitempty"`
	Alamat    string    `json:"alamat"  bson:"alamat,omitempty"`
	Umur      string    `json:"umur"  bson:"umur,omitempty"`
	NoHp      string    `json:"nohp"  bson:"nohp,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updateAt" bson:"update_at"`
}

type UserRequest struct {
	Nama   string `json:"nama"  bson:"nama,omitempty"`
	Alamat string `json:"alamat"  bson:"alamat,omitempty"`
	Umur   string `json:"umur"  bson:"umur,omitempty"`
	NoHp   string `json:"nohp"  bson:"nohp,omitempty"`
}
