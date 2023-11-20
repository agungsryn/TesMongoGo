package user

import (
	"be-go-mongo/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// var userCollection = db.connectDb().Database("goTest").Collection("users") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding Content-type
	var userRequest UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}

	user := User{
		Nama:   userRequest.Nama,
		Alamat: userRequest.Alamat,
		Umur:   userRequest.Umur,
		NoHp:   userRequest.NoHp,
		CreatedAt: time.Now(),
	}
	
	_, err = db.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", user)
	json.NewEncoder(w).Encode(user) // return the mongodb ID of generated document

}



