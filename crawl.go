package main

import (
    "context"
    "errors"
    "fmt"
    "io"
    "net/http"
    "time"
)

// MakeRequest performs an HTTP request with a timeout and error handling.
func MakeRequest(url string, timeout time.Duration) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return "", fmt.Errorf("failed to create request: %w", err)
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", fmt.Errorf("failed to execute request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("non-200 response status")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("failed to read response body: %w", err)
    }

    return string(body), nil
}

func main() {
    url := "http://example.com"
    timeout := 5 * time.Second

    response, err := MakeRequest(url, timeout)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }

    fmt.Printf("Response: %s\n", response)
}