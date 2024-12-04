package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchData() (string, int, error) {
	// Read environment variables
	url := os.Getenv("AOC_URL")
	cookieName := os.Getenv("COOKIE_NAME")
	cookieValue := os.Getenv("COOKIE_VALUE")

	if url == "" || cookieName == "" || cookieValue == "" {
		return "", 0, fmt.Errorf("AOC_URL, COOKIE_NAME, and COOKIE_VALUE environment variables must be set")
	}

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, fmt.Errorf("error creating request: %v", err)
	}

	// Add the cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	})

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, fmt.Errorf("error reading response body: %v", err)
	}

	return string(rawData), resp.StatusCode, nil
}

