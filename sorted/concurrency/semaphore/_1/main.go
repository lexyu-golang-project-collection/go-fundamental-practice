package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func fetchURL(url string, wg *sync.WaitGroup, sem chan struct{}) {
	defer func() {
		<-sem
		wg.Done()
	}()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error requesting %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("URL: %s, Status: %s\n", url, resp.Status)
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.reddit.com",
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 5)

	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{}
		go fetchURL(url, &wg, sem)
	}

	wg.Wait()
	close(sem)
}
