package main

import (
	"fmt"
	"github.com/PaavoReinikka/GoUtils/greetings"
)

func main() {
	msg := greetings.Hello("Paavo")
	fmt.Println(msg)
}
