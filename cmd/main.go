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
	"os/signal"
	"syscall"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
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
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: router,
	}
	// setup handle to goroutine
	go func() {
		// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
		log.Printf("Server listen on %s", PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// make channel for listen on os.Signal and setup notify signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	// setup withTimeout to preserve connection before close
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
	log.Println("Server exiting")
}
