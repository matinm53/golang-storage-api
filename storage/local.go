package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	os.MkdirAll(basePath, os.ModePerm)
	return &LocalStorage{BasePath: basePath}
}

func (s *LocalStorage) SaveFile(filename string, file io.Reader) error {
	fullPath := filepath.Join(s.BasePath, filename)
	out, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}

func (s *LocalStorage) ReadFile(filename string) (*os.File, error) {
	fullPath := filepath.Join(s.BasePath, filename)
	return os.Open(fullPath)
}
