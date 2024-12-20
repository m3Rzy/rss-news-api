package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rss-news-api/config"
	"rss-news-api/internal/models"
	"time"
)

// NewsAPIClient реализует интерфейс APIClient
type NewsAPIClient struct{}

// Реализация метода GetDataFromAPI
func (c *NewsAPIClient) GetDataFromAPI(inputValue string) error {
	// Формируем запрос
	requestURL := fmt.Sprintf(
		"https://newsapi.org/v2/everything?from=%s&q=%s&apiKey=%s",
		time.Now().Format(time.RFC3339),
		inputValue,
		config.GetConfig(),
	)

	// Выполняем запрос
	response, err := http.Get(requestURL)
	if err != nil {
		return fmt.Errorf("HTTP request error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: status %d", response.StatusCode)
	}

	// Чтение тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %v", err)
	}

	// Парсим JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("Failed to parse JSON: %v", err)
	}

	// Извлечение массива статей
	articles, ok := result["articles"].([]interface{})
	if !ok {
		return fmt.Errorf("Failed to convert articles to array")
	}

	// Обработка статей
	for _, article := range articles {
		articleMap, ok := article.(map[string]interface{})
		if !ok {
			return fmt.Errorf("Failed to convert article to map")
		}

		// Создаём экземпляр Article
		news := &models.Article{
			Author:      getStringFromMap(articleMap, "author"),
			Title:       getStringFromMap(articleMap, "title"),
			URL:         getStringFromMap(articleMap, "url"),
			PublishedAt: getStringFromMap(articleMap, "publishedAt"),
			Content:     getStringFromMap(articleMap, "content"),
			Description: getStringFromMap(articleMap, "description"),
		}

		// Выводим данные
		news.StructToString()
	}

	return nil
}
