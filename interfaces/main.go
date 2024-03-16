package main

import "fmt"

type bot interface {
	getGreeting() string
}

// any struct type that has implementations that matches int interface type it
// it belongs to that interface type.

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
