package interfaces

// APIClient описывает интерфейс для работы с различными API
type APIClient interface {
	GetDataFromAPI(inputValue string) error
}
