package user

import (
	"be-go-mongo/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Update Profile of User

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body User
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}

	paramsStrId := mux.Vars(r)["id"]

	objectId, err := primitive.ObjectIDFromHex(paramsStrId)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.M{"_id": objectId} // converting value to BSON type
	after := options.After            // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}

	update := bson.M{"$set": bson.M{
		"nama":   body.Nama,
		"alamat": body.Alamat,
		"umur":   body.Umur,
		"nohp":   body.NoHp,
		"updated_at" : time.Now(),
	}}
	updateResult := db.UserCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}