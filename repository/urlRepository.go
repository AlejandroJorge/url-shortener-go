package repository

import "github.com/AlejandroJorge/url-shortener-go/model"

type URLRepository interface {
	Save(model.URL) error
	GetOriginalURL(shortenedPath string) (string, error)
	GetPreviousStoredShortenedURL(originalURL string) (bool, string, error)
	UpdateRedirectionsCount(originalURL string) error
}
