package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    // Define flags
    verbose := flag.Bool("verbose", false, "Enable verbose output")
    quiet := flag.Bool("quiet", false, "Suppress all output")

    // Parse command-line flags
    flag.Parse()

    // Example HTTP client with timeout
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    // Sample request
    resp, err := client.Get("http://example.com")
    if err != nil {
        logError(err, *verbose, *quiet)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        logError(fmt.Errorf("failed with status: %s", resp.Status), *verbose, *quiet)
        return
    }

    if *verbose {
        fmt.Printf("Request successful: %s\n", resp.Status)
    }
}

// logError handles error logging based on the flags
func logError(err error, verbose, quiet bool) {
    if quiet {
        return
    }
    if verbose {
        log.Printf("ERROR: %v\n", err)
    } else {
        log.Println("An error occurred. Use -verbose for more details.")
    }
}