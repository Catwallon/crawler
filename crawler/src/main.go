package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly/v2"
)

type Page struct {
	Website     string
	Url         string
	Title       string
	Description string
	Keywords    string
	Score       uint
}

var mu sync.Mutex

func exit_err(msg string, err interface{}) {
	log.Fatalf("%s: %v", msg, err)
}

func connect_db() *sql.DB {
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

func startsWithHTTP(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func worker(db *sql.DB, rawURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	scrap_page(db, rawURL)
}

func scrap_page(db *sql.DB, rawURL string) {
	page := Page{}
	var urls []string
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Println("ERROR PARSING: ", rawURL)
		return
	}
	page.Website = parsedURL.Host
	page.Url = rawURL
	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		page.Title = e.Text
	})
	c.OnHTML("a", func(e *colly.HTMLElement) {
		if startsWithHTTP(e.Attr("href")) {
			urls = append(urls, e.Attr("href"))
		}
	})
	er := c.Visit(rawURL)
	if er != nil {
		log.Println("ERROR SCRAPING: ", rawURL)
	} else {
		index_page(db, page)
		log.Println("INDEXED: ", rawURL)
		next_url(db, urls)
	}
}

func next_url(db *sql.DB, urls []string) {
	var indexed bool
	for i := 0; i < len(urls); i++ {
		mu.Lock()
		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM pages WHERE url = ?)", urls[i]).Scan(&indexed)
		mu.Unlock()
		if err != nil {
			exit_err("Can't search into table", err)
		} else if !indexed {
			scrap_page(db, urls[i])
		}
	}
}

func index_page(db *sql.DB, page Page) {
	query := "INSERT INTO pages (website, url, title, description, keywords, score) VALUES (?, ?, ?, ?, ?, ?)"
	mu.Lock()
	_, err := db.Exec(query, page.Website, page.Url, page.Title, page.Description, page.Keywords, page.Score)
	mu.Unlock()
	if err != nil {
		//exit_err("Can't insert into table", err)
	}
}

func main() {
	var wg sync.WaitGroup
	db := connect_db()
	const nbWorkers = 5
	wg.Add(nbWorkers)
	for i := 0; i < nbWorkers; i++ {
		go worker(db, os.Getenv("CRAWLER_START_URL"), &wg)
	}
	wg.Wait()
	db.Close()
}
