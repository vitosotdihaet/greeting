package greeting

import "fmt"

func greet(name string) string {
	var msg string
	msg = fmt.Sprint("Hello, %v!", name)
	return msg
}

func farewell(name string) string {
	msg := fmt.Sprint("Goodbye, %v!", name)
	return msg
}
