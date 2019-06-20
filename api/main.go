package main

import (
	f "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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
	http.HandleFunc("/api/getshort/", getShortURL)

	f.Println("Server running on port 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
