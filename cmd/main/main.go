package main

import (
	"log"
	"rss-news-api/internal/interfaces"
	"rss-news-api/internal/services"
)

func main() {
	var client interfaces.APIClient = &services.NewsAPIClient{}

	if err := client.GetDataFromAPI("golang"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Successfully fetched data from NewsAPI")
}
