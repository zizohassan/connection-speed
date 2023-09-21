package main

import (
	"io"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	http.HandleFunc("/download", downloadTest)
	http.HandleFunc("/upload", uploadTest)
	http.HandleFunc("/ping", pingTest)
	http.ListenAndServe(":8080", nil)
}

func downloadTest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	data := make([]byte, 1024*1024) // 1 MB of data
	w.Write(data)
}

func uploadTest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Upload test complete"))
}

func pingTest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Pong"))
}
