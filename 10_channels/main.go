package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up")
	c <- "Yep its up"
}
