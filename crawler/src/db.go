package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mu sync.Mutex

func connectDB() *sql.DB {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	dsn := fmt.Sprintf("root:%s@tcp(mariadb:3306)/db", password)
	for i := 5; i > 0; i-- {
		db, err := sql.Open("mysql", dsn)
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
		return db
	}
	log.Fatal("Failed to connect to database, exiting")
	return nil
}

func indexPage(db *sql.DB, page Page) {
	query := "INSERT INTO pages (website, url, title, description, keywords, lang) VALUES (?, ?, ?, ?, ?, ?)"
	mu.Lock()
	_, err := db.Exec(query, page.Website, page.Url, page.Title, page.Description, page.Keywords, page.Lang)
	mu.Unlock()
	checkError("Can't insert into database", err, false)
}

func alreadyIndexed(db *sql.DB, URL string) bool {
	var indexed bool
	mu.Lock()
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM pages WHERE url = ?)", URL).Scan(&indexed)
	mu.Unlock()
	checkError("Connection to database lost", err, true)
	return indexed
}
