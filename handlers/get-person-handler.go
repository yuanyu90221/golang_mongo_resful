package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang_mongo_restful/dtos"
	"golang_mongo_restful/mongodb"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPersonHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person dtos.Person
	collection := mongodb.Client.Database(mongodb.MONGO_DATABASE).Collection(mongodb.MONGO_COLLECTION)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, dtos.Person{ID: id}).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(person)
}
