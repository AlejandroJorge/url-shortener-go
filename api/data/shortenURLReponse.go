package data

type ShortenURLResponse struct {
	ShortenedURL string `validate:"required min=1" json:"shortenedURL"`
}
