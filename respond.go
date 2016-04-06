package respond

import (
	"encoding/json"
	"net/http"
)

/*
Send a byte slice as a response.
*/
func ByteSlice(w http.ResponseWriter, message []byte, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write(message)
}

/*
Send a string as your response.
*/
func Simple(w http.ResponseWriter, message string, statusCode int) {
	ByteSlice(w, []byte(message), statusCode)
}

/*
Send a json-serializable object as the response.
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
