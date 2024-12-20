package main

import (
	"rss-news-api/internal/services"
)

func main() {
	// указать ключевое слово
	services.GetDataFromNewsApi("Russia")
}
