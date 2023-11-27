package customerror

import "errors"

var ErrNoShortenedRouteRegistered = errors.New("shortened path isn't registered")

var ErrShortenedRouteAlreadyExists = errors.New("shortened route already exists")

var ErrInvalidOriginalURL = errors.New("invalid original url")
