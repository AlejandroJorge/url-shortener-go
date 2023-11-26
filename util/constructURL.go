package util

import "fmt"

func ConstructURL(domain string, port uint, route string) string {
	return fmt.Sprintf("%s:%d/%s", domain, port, route)
}
