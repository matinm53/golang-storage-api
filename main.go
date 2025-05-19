package main

import (
	"log"
	"net/http"

	"github.com/matinm53/golang-storage-api/api"
)

func main() {
	http.HandleFunc("/upload", api.UploadHandler)
	http.HandleFunc("/file/", api.FileHandler)
	http.HandleFunc("/file-review/", api.FileReviewHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
