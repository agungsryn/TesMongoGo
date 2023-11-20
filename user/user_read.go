package user

import (
	"be-go-mongo/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                 //slice for multiple documents
	cur, err := db.UserCollection.Find(context.TODO(), bson.M{}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	var result primitive.M //  an unordered representation of a BSON document which is a Map
	paramsStrId := mux.Vars(r)["id"]

	objectId, err := primitive.ObjectIDFromHex(paramsStrId)
	if err != nil {
		log.Println("Invalid id")
	}

	fmt.Println("isi", objectId)

	err = db.UserCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document
}