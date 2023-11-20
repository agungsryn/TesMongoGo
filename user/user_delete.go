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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Delete Profile of User

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	paramsStrId := mux.Vars(r)["id"] //get Parameter value as string

	objectId, err := primitive.ObjectIDFromHex(paramsStrId) // convert params to mongodb Hex ID
	if err != nil {
		log.Println("Invalid id")
	}

	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := db.UserCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}