package main

// import (
// 	"database/sql"
// 	f "fmt"
// 	"io/ioutil"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"net/url"

// 	_ "github.com/mattn/go-sqlite3"
// )

import (
	"database/sql"
	f "fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	_ "github.com/mattn/go-sqlite3"
)

type urlmap struct {
	short string
	long  string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var db *sql.DB

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

	// Generate short url
	shortURL := getRandString(5)

	// Check if db alread has short url
	sqlCheck := "select maps.Short from maps where maps.Short = '" + shortURL + "'"
	shortInDB, err := db.Query(sqlCheck)
	if err != nil {
		log.Fatal(err)
	}
	defer shortInDB.Close()

	var shortStatus string // Status of the short url; if it's already in db or not
	shortInDB.Scan(&shortStatus)
	if shortStatus == "" {
		sqlInsert := "insert into maps values ('" + shortURL + "'," + "'" + decodedURL + "')"
		_, err := db.Exec(sqlInsert)
		if err != nil {
			log.Fatal(err)
			w.Write([]byte("null"))
		} else {
			w.Write([]byte(shortURL))
		}
	} else {
		w.Write([]byte("null"))
	}
}

func main() {
	http.HandleFunc("/api/getshort/", getShortURL)

	var err error
	db, err = sql.Open("sqlite3", "./urlmap.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		f.Println(err)
	}

	f.Println("Server running on port 9000")

	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
