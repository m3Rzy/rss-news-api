package models

import "fmt"

type Article struct {
	Author      string
	Title       string
	URL         string
	PublishedAt string
	Content     string
	Description string
}

func (a *Article) String() string {
	return fmt.Sprintf(
		"Автор: %s\nЗаголовок: %s\nПуть до ресурса: %s\nОпубликован: %s\nКонтент: %s\nОписание: %s\n\n",
		a.Author, a.Title, a.URL, a.PublishedAt, a.Content, a.Description,
	)
}

func (a *Article) StructToString() {
	fmt.Print(a.String())
}
