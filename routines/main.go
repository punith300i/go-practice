package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinks(link, c)
	}

	for l := range c {

		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLinks(l, c)
		}(l)

	}
}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- " Might be down "
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
