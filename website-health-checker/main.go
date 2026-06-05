package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Url    string
	Status string
	Err    error
}

func checkWebsiteUrl(url string, ch chan Result) {
	res, err := http.Get(url)

	if err != nil {
		// fmt.Println(url, "is down")

		// sending data into channel
		ch <- Result{
			Url:    url,
			Status: "is down",
			Err:    err,
		}
		return
	}

	defer res.Body.Close()

	ch <- Result{
		Url:    url,
		Status: "is up and running",
		Err:    nil,
	}

}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://wrong-url-test.com",
	}

	// channel
	ch := make(chan Result)

	start := time.Now()

	for _, url := range urls {
		go checkWebsiteUrl(url, ch)
	}

	for range urls {
		result := <-ch // Blocking call, waits for a result from the channel

		// if result.Err != nil {
		// 	fmt.Println(result.Url, result.Status, "Error:", result.Err)
		// 	continue
		// }

		fmt.Println("Resut", result)
	}

	fmt.Println("time taken", time.Since(start))
	fmt.Println("All url's checked successfully")
}
