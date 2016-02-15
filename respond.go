package respond

import (
	"encoding/json"
	"net/http"
)

/*
Send a string as your response.
*/
func Simple(w http.ResponseWriter, message String, statusCode int) {
	w.WriteHeader(statusCode)
	response.Write([]byte(message))
}

/*
Send a json-serializable object as the response.
*/
func Json(w http.ResponseWriter, message interface{}, statusCode int) error {
	jsonAsString, err := json.Marshal(message)

	if err != nil {
		return err
	}

	Simple(w, jsonAsString, statusCode)
	return nil
}
