package greeting

import "fmt"

func greet(name string) string {
	var msg string = fmt.Sprintf("Hello, %v!", name)
	return msg
}

func farewell(name string) string {
	msg := fmt.Sprintf("Goodbye, %v!", name)
	return msg
}
