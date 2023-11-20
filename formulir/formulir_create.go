package formulir

import (
	"be-go-mongo/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CreateFormulir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var req FormulirRequest
	err := json.NewDecoder(r.Body).Decode(&req) // storing in person variable of type user
	if err != nil {
		fmt.Println(err)
	}

	// objectId, err := primitive.ObjectIDFromHex("6558bfa4d9efcefccfa1488c")
	// if err != nil {
	// 	log.Println("Invalid id")
	// }
	// dateString := "2021-11-22"
	fmt.Println("req.ValidThru",req.ValidThru)
	date, _ := time.Parse("2006-01-02", req.ValidThru)
	formulir := Formulir{
		UserID:      req.UserID,
		Description: req.Description,
		CreatedAt:   time.Now(),
		ValidThru:   date,
	}

	fmt.Println("before: ", formulir)

	_, err = db.FormulirCollection.InsertOne(context.TODO(), formulir)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(formulir) // return the mongodb ID of generated document

}
