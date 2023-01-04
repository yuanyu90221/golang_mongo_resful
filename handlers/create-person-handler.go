package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_mongo_restful/dtos"
	"golang_mongo_restful/mongodb"
	"net/http"
	"time"
)

func CreatePersonHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	var person dtos.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	fmt.Println(person, mongodb.MONGO_COLLECTION, mongodb.MONGO_DATABASE)
	collection := mongodb.Client.Database(mongodb.MONGO_DATABASE).Collection(mongodb.MONGO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(w).Encode(result)
}
