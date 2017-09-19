package main

import (
	"net/http"
	"os"
	"log"
)

func main() {
	host := env("HOST", "")
	port := env("PORT", "8080")
	folder := env("FOLDER", "/web") + "/"
	file := env("FILENAME", "file.pdf")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, folder+file)
	})
	http.ListenAndServe(host+":"+port, nil)
}

func env(key, fallback string) string {
	if value := os.Getenv(key); 0 < len(value) {
		return value
	}
	return fallback
}
