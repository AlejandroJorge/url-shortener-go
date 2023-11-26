package data

type ShortenURLRequest struct {
	OriginalURL string ` validate:"required min=1" json:"originalURL"`
}
