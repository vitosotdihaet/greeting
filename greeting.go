package greeting

import "fmt"

func Greet(name string) string {
	var msg string = fmt.Sprintf("Hello, %v!", name)
	return msg
}

func Farewell(name string) string {
	var msg string = fmt.Sprintf("Goodbye, %v!", name)
	return msg
}
