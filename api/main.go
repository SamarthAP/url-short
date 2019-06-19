package main

import (
	f "fmt"
	"log"
	"net/http"
)

func main() {
	f.Println("Server running on port 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
