package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rss-news-api/config"
	"time"
)

func GetDataFromNewsApi(inputValue string) {
	requestURL := fmt.Sprintf(
		"https://newsapi.org/v2/everything?from=%s&q=%s&apiKey=%s",
		time.Date(2024, time.December, 19, 10, 10, 10, 10, time.UTC),
		inputValue,
		config.GetConfig(),
	)

	response, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Error HTTP-request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка API: статус %d", response.StatusCode)
	}

	// Чтение тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения тела ответа: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	// Извлечение массива articles
	articles, ok := result["articles"].([]interface{})
	if !ok {
		log.Fatalf("Ошибка преобразования articles в массив")
	}

	// Проход по всем статьям и вывод авторов и заголовков
	for _, article := range articles {
		articleMap, ok := article.(map[string]interface{})
		if !ok {
			log.Fatalf("Ошибка преобразования статьи в объект")
		}

		// Извлечение авторов и заголовков
		author, _ := articleMap["author"].(string)
		title, _ := articleMap["title"].(string)
		url, _ := articleMap["url"].(string)
		publishedAt, _ := articleMap["publishedAt"].(string)
		content, _ := articleMap["content"].(string)
		description, _ := articleMap["description"].(string)

		// Вывод
		fmt.Printf("Автор: %s\n", author)
		fmt.Printf("Заголовок: %s\n", title)
		fmt.Printf("Путь до ресурса: %s\n", url)
		fmt.Printf("Опубликован: %s\n", publishedAt)
		fmt.Printf("Контент: %s\n", content)
		fmt.Printf("Описание: %s\n\n", description)
	}
}
