package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var start = time.Now()

	// wg.Add(1)
	// go uploadFile()
	wg.Go(uploadFile)

	// wg.Add(1)
	// go SaveToDB()
	wg.Go(SaveToDB)

	// wg.Add(1)
	// go sendEmail()
	wg.Go(sendEmail)

	wg.Wait()

	fmt.Println("time taken", time.Since(start))
}

func uploadFile() {
	// defer wg.Done()
	fmt.Println("Uploading file")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done")
	// wg.Add(-1)
	// wg.Done()
}

func SaveToDB() {
	fmt.Println("Saving to db")
	time.Sleep(1 * time.Second)
	fmt.Println("Saved to db")
	// wg.Add(-1)
	// wg.Done()
}

func sendEmail() {
	fmt.Println("sending email")
	time.Sleep(2 * time.Second)
	fmt.Println("email sent")
	// wg.Add(-1)
	// wg.Done()
}
