package services

// Вспомогательная функция для безопасного извлечения строки из map
func getStringFromMap(data map[string]interface{}, key string) string {
	if value, ok := data[key].(string); ok {
		return value
	}
	return ""
}
