package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseObj struct {
	Collection *mongo.Collection
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome Page")
}

func (db DatabaseObj) rootEndpoint(writer http.ResponseWriter, request *http.Request) {
	var data interface{}
	byteData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal("IO Error: ", err)
	}

	err = json.Unmarshal(byteData, &data)
	if err != nil {
		log.Fatal("JSON Error: ", err)
	}

	status, err := db.Collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal("Insert Error: ", err)
	}
	fmt.Fprintf(writer, "\n\n%+v", status)

	fmt.Fprintf(writer, "\n\n%+v", data)
}

func main() {
	mongoservice := os.Getenv("MONGO_SERVICE_NAME")
	port := ":" + os.Getenv("API_PORT")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoservice + ":27017"))
	if err != nil {
		log.Fatal("Mongo Error: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Connection Error: ", err)
	}
	defer client.Disconnect(ctx)

	databaseObj := DatabaseObj{
		Collection: client.Database("Mydatabase").Collection("Mycollection"),
	}

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/", databaseObj.rootEndpoint).Methods("POST")
	fmt.Println("RESTful-API-Go-Endpoint Hosted on port " + port)
	log.Fatal("HTTP Error: ", http.ListenAndServe(port, router))
}
