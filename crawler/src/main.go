package main

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Website     string
	Url         string
	Title       string
	Description string
	Keywords    []byte
	Lang        string
}

func worker(db *sql.DB, URL string, wg *sync.WaitGroup) {
	defer wg.Done()
	scrapPage(db, URL)
}

func main() {
	var wg sync.WaitGroup
	loadStopwords()
	db := connectDB()
	const nbWorkers = 5
	wg.Add(nbWorkers)
	for i := 0; i < nbWorkers; i++ {
		go worker(db, os.Getenv("CRAWLER_START_URL"), &wg)
	}
	wg.Wait()
	db.Close()
}
