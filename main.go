package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define the /hello endpoint
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        // Set the response header
        w.Header().Set("Content-Type", "text/plain")
        // Write the response
        fmt.Fprintln(w, "hello")
    })

    // Start the web server on port 8080
    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
