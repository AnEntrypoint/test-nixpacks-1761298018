package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Nixpacks Test App!\nDeployed via Coolify CLI\nTimestamp: %s\n", time.Now().Format("2006-01-02 15:04:05"))
    })
    
    fmt.Printf("Server starting on port 3000\n")
    http.ListenAndServe(":3000", nil)
}
