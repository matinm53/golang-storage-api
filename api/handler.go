package api

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"

	"github.com/matinm53/golang-storage-api/storage"
)

var store = storage.NewLocalStorage("./uploads")

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

	err = store.SaveFile(header.Filename, file)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Uploaded successfully"))
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	filename := path.Base(strings.TrimPrefix(r.URL.Path, "/file/"))
	file, err := store.ReadFile(filename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, filename, info.ModTime(), file)
}

func FileReviewHandler(w http.ResponseWriter, r *http.Request) {
	filename := path.Base(strings.TrimPrefix(r.URL.Path, "/file-review/"))
	file, err := store.ReadFile(filename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		http.Error(w, "Error reading file info", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"name":         info.Name(),
		"size":         info.Size(),
		"modifiedTime": info.ModTime(),
	})
}
