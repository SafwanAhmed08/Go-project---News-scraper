package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Article struct {
    Links []string `json:"links"`
}

var mutex sync.Mutex
var articles []Article

func main() {

	// User inputs the page name
	var keyword string
	fmt.Print("Enter the name of the News link page: ")
	fmt.Scanln(&keyword)

	var wg sync.WaitGroup
    wg.Add(2)
	
	go func(){
		defer wg.Done()
		scrape("https://timesofindia.indiatimes.com/topic/",keyword)
	}()

	go func() {
        defer wg.Done()
        scrape("https://www.ndtv.com/search?searchtext=", keyword)
    }()
	
	wg.Wait()

	exportToJSON(articles)
}

func scrape(newssite string, keyword string){
	c := colly.NewCollector()
	var selector string

	if(newssite == "https://timesofindia.indiatimes.com/topic/"){
		selector = ".crmK8"
	}else{
		selector = ".src_tab-cnt"
	}
	
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		article := Article{Links: links}

		mutex.Lock()
		articles = append(articles, article)
		mutex.Unlock()
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
		os.Exit(1) // Exit or handle according to your need
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraping finished.")
	})

	

	url := newssite + strings.ReplaceAll(keyword, " ", "-")

	fmt.Println("Visiting", url)

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		log.Fatal("Error visiting page:", err)
	}
	c.Wait()

	// // After scraping, export data
	exportToJSON(articles)

}

func exportToJSON(articles []Article) {
	// Create a new JSON file
	jsonFile, err := os.Create("articles.json")
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
	}
	defer jsonFile.Close()

	// Marshal articles to JSON
	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	// Write JSON data to file
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing JSON data to file:", err)
	}
}