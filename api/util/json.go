package util

import (
	"encoding/json"
	"net/http"

	"github.com/AlejandroJorge/url-shortener-go/data"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteResponseHeaders(w http.ResponseWriter, response data.WebResponse) {
	w.WriteHeader(int(response.Code))
	w.Header().Set("Status", response.Status)
}

func WriteResponseBody(w http.ResponseWriter, response data.WebResponse) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response.Data)
	PanicIfError(err)
}
