# Golang Storage API

A simple Go package and server to upload and read files using HTTP.

## Endpoints

- `POST /upload` — Form file field: `file`
- `GET /file/{filename}` — Returns the uploaded file

## Run

```bash
go run main.go