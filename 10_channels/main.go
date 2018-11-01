package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		// use go keyword only in front of dunction calls
		// channels are used to communicate between go routines
		// Recieving messages over channel is a blocking process
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	//Alterate syntax
	for l := range c {
		// time.Sleep(time.Second)
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
