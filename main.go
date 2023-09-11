package main

import (
    "os"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    podName := os.Getenv("POD_NAME")
    w.Write([]byte("POD_NAME: " + podName))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
