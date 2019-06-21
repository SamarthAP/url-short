package main

import (
	f "fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

type urlmap struct {
	short string
	long  string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func getRandString(length int) string {
	str := make([]rune, length)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
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

	// shortURL := getRandString(5)

	// Generate 5 letter random string from A-Za-z0-9
	// Check if string exists in db
	// If not we use that string and create a db entry for it
	// Send the string back as the response

	w.Write([]byte(decodedURL))
}

func main() {
	http.HandleFunc("/api/getshort/", getShortURL)

	f.Println("Server running on port 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
