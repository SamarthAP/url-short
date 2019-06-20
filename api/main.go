package main

import (
	"context"
	f "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://dbadmin:dbadmin@urlmaps-74jtc.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	f.Println("Connected to database")
}

func getShortURL(w http.ResponseWriter, r *http.Request) {
	longurl, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading body:", err)
		http.Error(w, "can't ready body", http.StatusBadRequest)
		return
	}
	decodedURL, err := url.QueryUnescape(string(longurl))
	if err != nil {
		log.Fatal("Error decoding url:", err)
		http.Error(w, "can't decode url", http.StatusNotAcceptable)
		return
	}
	w.Write([]byte(decodedURL))
}

func main() {
	connectMongo()

	http.HandleFunc("/api/getshort/", getShortURL)

	f.Println("Server running on port 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
