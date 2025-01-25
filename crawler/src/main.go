package main

import (
	"database/sql"
	"errors"
	"os"
	"strconv"
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
	nbWorkers, err := strconv.Atoi(os.Getenv("CRAWLER_NB_WORKERS"))
	checkError("Invalid CRAWLER_NB_WORKERS value", err, true)
	if nbWorkers < 1 {
		checkError("Invalid CRAWLER_NB_WORKERS value", errors.New("lower than 1"), true)
	}
	wg.Add(nbWorkers)
	for i := 0; i < nbWorkers; i++ {
		go worker(db, os.Getenv("CRAWLER_START_URL"), &wg)
	}
	wg.Wait()
	db.Close()
}
