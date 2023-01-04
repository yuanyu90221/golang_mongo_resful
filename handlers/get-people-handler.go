package handlers

import (
	"context"
	"encoding/json"
	"golang_mongo_restful/dtos"
	"golang_mongo_restful/mongodb"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPeopleHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	var people []dtos.Person
	collection := mongodb.Client.Database(mongodb.MONGO_DATABASE).Collection(mongodb.MONGO_COLLECTION)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person dtos.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(&people)
}
