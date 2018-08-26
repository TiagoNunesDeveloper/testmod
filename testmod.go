package testmod

import "fmt"

// Hi retorna uma saudação
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
