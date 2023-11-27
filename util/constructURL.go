package util

import "fmt"

func ConstructURL(domainWithPort string, route string) string {
	return fmt.Sprintf("%s/%s", domainWithPort, route)
}
