package respond

import (
	"encoding/json"
	"net/http"
)

/*
ByteSlice send a byte slice as a response.
*/
func ByteSlice(w http.ResponseWriter, message []byte, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write(message)
}

/*
Simple sends a string as your response.
*/
func Simple(w http.ResponseWriter, message string, statusCode int) {
	ByteSlice(w, []byte(message), statusCode)
}

/*
JSON sends a json-serializable object as the response.
*/
func JSON(w http.ResponseWriter, message interface{}, statusCode int) error {
	jsonAsByteSlice, err := json.Marshal(message)

	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	ByteSlice(w, jsonAsByteSlice, statusCode)
	return nil
}

/*
HTML sends an html string as the response.
*/
func HTML(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Add("Content-Type", "text/html")
	Simple(w, message, statusCode)
}
