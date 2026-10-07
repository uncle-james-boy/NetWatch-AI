package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 1. THE BLUEPRINT (Struct)
// We define what our telemetry data will look like.
// The text in backticks (like `json:"target"`) tells Go what to name the fields in JSON.
type HTTPResult struct {
	Target     string  `json:"target"`
	StatusCode int     `json:"status_code"`
	LatencyMs  float64 `json:"latency_ms"`
	Timestamp  string  `json:"timestamp"`
	Error      string  `json:"error,omitempty"` // "omitempty" means don't print this if there is no error
}

func main() {
	url := "http://www.google.com"

	// 2. MEASURE THE TIME
	startTime := time.Now()
	resp, err := http.Get(url)
	endTime := time.Now()

	// Calculate latency in milliseconds (float64)
	latencyMs := float64(endTime.Sub(startTime).Milliseconds())

	// 3. FILL OUT THE BLUEPRINT
	result := HTTPResult{
		Target:    url,
		LatencyMs: latencyMs,
		Timestamp: time.Now().Format(time.RFC3339), // Standard cloud time format
	}

	// 4. ERROR HANDLING
	if err != nil {
		result.Error = err.Error()
		result.StatusCode = 0
	} else {
		result.StatusCode = resp.StatusCode
	}

	// 5. CONVERT TO JSON (This is called "Marshaling")
	jsonData, jsonErr := json.MarshalIndent(result, "", "  ")
	if jsonErr != nil {
		fmt.Println("Failed to create JSON:", jsonErr)
		return
	}

	// 6. PRINT THE JSON
	fmt.Println(string(jsonData))
}
