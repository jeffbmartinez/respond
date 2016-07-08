package respond

import (
	"bytes"
	"encoding/json"
	"html/template"
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
String sends a string as your response.
*/
func String(w http.ResponseWriter, message string, statusCode int) {
	ByteSlice(w, []byte(message), statusCode)
}

/*
Simple calls String, populating message with the equivalent reason
phrase as returned by http.StatusText, based on the
statusCode.

For example, Status(w, http.StatusBadRequest) is the equivalent
of String(w, "Bad Request", http.StatusBadRequest)

If the statusCode is not recognized an empty string is used.
*/
func Simple(w http.ResponseWriter, statusCode int) {
	String(w, http.StatusText(statusCode), statusCode)
}

/*
JSON sends a json-serializable object as the response and sets
"Content-Type" header to "application/json".
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
HTML sends an html string as the response and sets
"Content-Type" header to "text/html".
*/
func HTML(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Add("Content-Type", "text/html")
	String(w, message, statusCode)
}

/*
HTMLTemplate renders an html template and sends an html string as the response. Will return an error
if the template was not found or otherwise had a problem. In the case of an error,
the http.ResponseWriter object will not be written to or modified and it is the developer's
responsibility to take appropriate action, such as respond.Simple(w, http.StatusInternalServerError)
*/
func HTMLTemplate(w http.ResponseWriter, templateName string, data interface{}, statusCode int) error {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		return err
	}

	HTML(w, b.String(), statusCode)
	return nil
}
