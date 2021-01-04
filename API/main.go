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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseObject struct {
	Collection	*mongo.Collection
	Context		context.Context
}

func (db DatabaseObject) GetEndpoint(writer http.ResponseWriter, request *http.Request) {
	var data []map[string]interface{}

	cursor, err := db.Collection.Find(db.Context, bson.M{})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(db.Context)

	for cursor.Next(db.Context) {
		var doc map[string]interface{}
		cursor.Decode(&doc)
		data = append(data, doc)
	}

	if err := cursor.Err(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
	}

	json.NewEncoder(writer).Encode(data)
}

func (db DatabaseObject) PostEndpoint(writer http.ResponseWriter, request *http.Request) {
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
	
	json.NewEncoder(writer).Encode(status)
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

	dbo := DatabaseObject{
		Collection: client.Database("Mydatabase").Collection("Mycollection"),
	}

	router := mux.NewRouter()
	router.HandleFunc("/", dbo.GetEndpoint).Methods("GET")
	router.HandleFunc("/", dbo.PostEndpoint).Methods("POST")
	fmt.Println("RESTful-API-Go-Endpoint Hosted on port " + port)
	log.Fatal("HTTP Error: ", http.ListenAndServe(port, router))
}
