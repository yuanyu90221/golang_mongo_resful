package handlers

import (
	"context"
	"encoding/json"
	"golang_mongo_restful/dtos"
	"golang_mongo_restful/mongodb"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePersonHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := mongodb.Client.Database(mongodb.MONGO_DATABASE).Collection(mongodb.MONGO_COLLECTION)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.DeleteOne(ctx, dtos.Person{ID: id})
	json.NewEncoder(w).Encode(result)
}
