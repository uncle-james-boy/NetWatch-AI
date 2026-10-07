package main

import (
	"fmt"
	"net/http"
	"time"
)

// This is the main function. When you run the program, this is where it starts.
func main() {
	// The URL we want to test
	url := "http://www.google.com"

	// We record the exact time right before we make the request
	startTime := time.Now()

	// We make a "GET" request to the URL. 
	// Think of this like your browser asking the server for the webpage.
	resp, err := http.Get(url)

	// We record the exact time right after the server replies
	endTime := time.Now()

	// Error handling: If the website is down or we have no internet, 'err' will not be empty.
	if err != nil {
		fmt.Println("❌ Network Error:", err)
		return // Stop the program here if there is an error
	}

	// Calculate how long the trip took
	latency := endTime.Sub(startTime)

	// Print the results to the terminal
	fmt.Printf("✅ Successfully reached %s\n", url)
	fmt.Printf("📊 Status Code: %d\n", resp.StatusCode)
	fmt.Printf("⏱️ Latency: %v\n", latency)
}