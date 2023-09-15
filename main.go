package main

import (
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	podName := os.Getenv("POD_NAME")
	url := r.URL.String()
	host := r.Host
	path := r.URL.Path
	scheme := r.URL.Scheme

	w.Write([]byte("POD_NAME: " + podName + "\nURL: " + url + "\nHOST: " + host + "\nPATH: " + path + "\nSCHEME: " + scheme))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
