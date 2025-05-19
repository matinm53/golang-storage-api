package api

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.MkdirAll(uploadDir, os.ModePerm)
	out, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	io.Copy(out, file)
	w.Write([]byte("Uploaded successfully"))
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path)
	filePath := filepath.Join(uploadDir, fileName)

	http.ServeFile(w, r, filePath)
}
