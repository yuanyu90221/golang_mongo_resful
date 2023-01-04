package main

import (
	"context"
	"fmt"
	dtos "golang_mongo_restful/dtos"
	"golang_mongo_restful/handlers"
	"golang_mongo_restful/mongodb"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var People []dtos.Person

func main() {
	// read config from .env
	godotenv.Load()
	PORT := os.Getenv("PORT")
	MONGO_URI := os.Getenv("MONGO_URI")
	// greeting
	log.Println("starting server ...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// setup mongo connect
	mongodb.SetUpConnect(MONGO_URI, ctx)
	mongodb.MONGO_COLLECTION = os.Getenv("MONGO_COLLECTION")
	mongodb.MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	// create router
	router := mux.NewRouter()
	router.HandleFunc("/people", handlers.GetPeopleHandler).Methods("GET")
	router.HandleFunc("/person/{id}", handlers.GetPersonHandler).Methods("GET")
	router.HandleFunc("/person", handlers.CreatePersonHandler).Methods("POST")
	router.HandleFunc("/person/{id}", handlers.DeletePersonHandler).Methods("DELETE")
	fmt.Printf("server listen on %s\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}
