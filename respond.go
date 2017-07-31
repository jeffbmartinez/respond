package respond

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
)

/*
ByteSlice send a byte slice as a response.

Returns the result of calling Write on the http.ResponseWriter interface, this
is most likely the number of bytes written and any error.
*/
func ByteSlice(w http.ResponseWriter, statusCode int, message []byte) (int, error) {
	w.WriteHeader(statusCode)
	return w.Write(message)
}

/*
String sends a string as your response.
*/
func String(w http.ResponseWriter, statusCode int, message string) (int, error) {
	return ByteSlice(w, statusCode, []byte(message))
}

/*
Simple calls String, populating message with the equivalent reason
phrase as returned by http.StatusText, based on the
statusCode.

For example, Status(w, http.StatusBadRequest) is the equivalent
of String(w, http.StatusBadRequest, "Bad Request")

If the statusCode is not recognized by http.StatusText an empty string is used.
*/
func Simple(w http.ResponseWriter, statusCode int) (int, error) {
	return String(w, statusCode, http.StatusText(statusCode))
}

/*
JSON sends a json-serializable object as the response and sets
"Content-Type" header to "application/json".

If there is an error it is returned and the Content-Type is not set, the
http.ResponseWriter object is not written to or modified.
*/
func JSON(w http.ResponseWriter, statusCode int, message interface{}) (int, error) {
	jsonAsByteSlice, err := json.Marshal(message)

	if err != nil {
		return 0, err
	}

	w.Header().Add("Content-Type", "application/json")
	return ByteSlice(w, statusCode, jsonAsByteSlice)
}

/*
HTML sends an html string as the response and sets
"Content-Type" header to "text/html".
*/
func HTML(w http.ResponseWriter, statusCode int, message string) (int, error) {
	w.Header().Add("Content-Type", "text/html")
	return String(w, statusCode, message)
}

/*
HTMLTemplate renders an html template and sends an html string as the response.

It will return an error if the template was not found or otherwise had a problem.

In the case of an error, the http.ResponseWriter object will not be written to or
modified. The developer should take appropriate action, such as
respond.Simple(w, http.StatusInternalServerError)
*/
func HTMLTemplate(w http.ResponseWriter, statusCode int, templateName string, data interface{}) (int, error) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		return 0, err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		return 0, err
	}

	return HTML(w, statusCode, b.String())
}
