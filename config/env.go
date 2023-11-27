package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlejandroJorge/url-shortener-go/util"
)

func GetPortString() string {
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	util.PanicIfError(err)

	return fmt.Sprintf(":%d", port)
}
