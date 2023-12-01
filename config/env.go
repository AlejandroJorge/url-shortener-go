package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlejandroJorge/url-shortener-go/util"
)

func GetPortString() string {
	portString := os.Getenv("PORT")
	if portString == "" {
		return fmt.Sprintf(":3000")
	}

	port, err := strconv.Atoi(portString)
	util.PanicIfError(err)

	return fmt.Sprintf(":%d", port)
}
