package server

import "net/http"

func setDefaultHeaders(writer http.ResponseWriter) http.ResponseWriter {
	writer.Header().Set("Content-Type", "application/json")
	return writer
}
