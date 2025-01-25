package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type Page struct {
	Website     string `json:"Website"`
	Url         string `json:"Url"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Keywords    string `json:"Keywords"`
	Score       int    `json:"Score"`
}

var db *sql.DB

func connect_db() {
	var err error
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	dsn := fmt.Sprintf("root:%s@tcp(mariadb:3306)/db", password)
	for i := 5; i > 0; i-- {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Println("Can't connect database", err)
			log.Println("Retry in 10 sec... ", i-1, " try remaining")
			time.Sleep(10 * time.Second)
			continue
		}
		err = db.Ping()
		if err != nil {
			log.Println("Can't connect database", err)
			log.Println("Retry in 10 sec... ", i-1, " try remaining")
			time.Sleep(10 * time.Second)
			db.Close()
			continue
		}
		log.Println("Successfully connected to database")
		break
	}
}

func main() {
	connect_db()
	mux := http.NewServeMux()

	mux.HandleFunc("/search", search)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)
	log.Fatal(http.ListenAndServeTLS(":"+os.Getenv("API_PORT"), "/cert/certificat.crt", "/cert/private.key", handler))
}

func search(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	search := queryParams.Get("query")
	log.Printf("%s %s", r.Method, r.URL.String())
	query := "SELECT website, url, title, description FROM pages WHERE title LIKE ? OR description LIKE ? OR keywords LIKE ?"
	rows, err := db.Query(query, "%"+search+"%", "%"+search+"%", "%"+search+"%")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer rows.Close()
	var pages []Page

	for rows.Next() {
		var page Page
		err := rows.Scan(
			&page.Website,
			&page.Url,
			&page.Title,
			&page.Description,
		)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		pages = append(pages, page)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}
