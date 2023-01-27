package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
	}

	t := time.Now()
	for _, link := range links {
		checkLink(link)
	}
	fmt.Println(">>>>>", time.Since(t))
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down")
		return
	}
	fmt.Println(link, " is up")
}
