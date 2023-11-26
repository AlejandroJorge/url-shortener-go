package repository

import "github.com/AlejandroJorge/url-shortener-go/model"

type URLRepository interface {
	Save(model.URL)
	GetOriginalURL(shortenedPath string) string
}
