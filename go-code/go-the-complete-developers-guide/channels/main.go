package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://twitter.com",
		"http://instagram.com",
	}

	// 1) Makes a channel
	c := make(chan string)

	for _, link := range links {
		// 2) Send links into the channel
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		// 3) Wait until something is received from the channel
		// This is a blocking call and will wait until something
		// is recied on the channel.
		fmt.Println(<-c)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, "- might be down!")
		// Send message into the channel
		c <- "Might be down I think!"
		return
	}
	fmt.Println(link, "- is up!")
	// Send message into the channel
	c <- "Yep it's up!"
}
