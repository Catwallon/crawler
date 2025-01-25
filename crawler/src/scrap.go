package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly/v2"
)

func findKeywords(words []string, lang string, nb int) []string {
	words = removeStopwords(words, lang)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	type wordFreq struct {
		Word  string
		Count int
	}
	frequencies := make([]wordFreq, 0, len(wordCount))
	for word, count := range wordCount {
		frequencies = append(frequencies, wordFreq{word, count})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Count > frequencies[j].Count
	})

	keywords := make([]string, 0, nb)
	for i := 0; i < len(frequencies) && i < nb; i++ {
		keywords = append(keywords, frequencies[i].Word)
	}
	return keywords
}

func sortURLs(sourceURL *url.URL, URLs []*url.URL) []*url.URL {
	sort.Slice(URLs, func(i, j int) bool {
		return URLs[i].Host != getDomain(sourceURL.Host) && getDomain(URLs[j].Host) == getDomain(sourceURL.Host)
	})
	sort.Slice(URLs, func(i, j int) bool {
		return URLs[i].Host != sourceURL.Host && URLs[j].Host == sourceURL.Host
	})
	return URLs
}

func scrapPage(db *sql.DB, URL string) {
	page := Page{}
	splitWordRegex := regexp.MustCompile(`[a-zA-ZÀ-ÿ]+`)

	var foundURLs []*url.URL
	var foundWords []string

	parsedURL, err := url.Parse(URL)
	checkError("Can't parse URL \""+URL+"\"", err, false)

	page.Website = parsedURL.Host
	page.Url = URL

	c := colly.NewCollector()
	c.OnHTML("html", func(e *colly.HTMLElement) {
		page.Lang = e.Attr("lang")
	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		page.Title = e.Text
	})
	c.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
		page.Description = e.Attr("content")
	})
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.DOM.Find("p, h1, h2, h3").Each(func(_ int, el *goquery.Selection) {
			text := el.Text()
			foundWords = append(foundWords, splitWordRegex.FindAllString(strings.ToLower(text), -1)...)
		})
	})
	c.OnHTML("a", func(e *colly.HTMLElement) {
		foundURL := e.Attr("href")
		parsedFoundURL, err := url.Parse(foundURL)
		if checkError("Can't parse URL \""+foundURL+"\"", err, false) && parsedFoundURL.Scheme == "https" {
			parsedFoundURL.RawQuery = ""
			foundURLs = append(foundURLs, parsedFoundURL)
		}
	})
	err = c.Visit(URL)
	if checkError("Can't scrap \""+URL+"\"", err, false) {
		if page.Lang == "" {
			page.Lang = "en"
		}
		keywords := findKeywords(foundWords, page.Lang, 10)
		page.Keywords, err = json.Marshal(keywords)
		if checkError("Can't marshal {"+strings.Join(keywords, ", ")+"}", err, false) && !alreadyIndexed(db, URL) {
			indexPage(db, page)
			log.Println("INDEXED:", URL)
		}
	}
	foundURLs = sortURLs(parsedURL, foundURLs)
	for _, url := range foundURLs {
		if !alreadyIndexed(db, url.String()) {
			scrapPage(db, url.String())
		}
	}
}
