package greetings

import "fmt"

func Hello(name string) (msg string) {
	msg = fmt.Sprintf("Hello %s!\n", name)
	return
}
