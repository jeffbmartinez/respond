package respond

import (
	"encoding/json"
	"net/http"
)

/*
ReasonPhrases is a map of status codes to reason phrases
as recommended by RFC 2616 section 6.1.1
https://www.ietf.org/rfc/rfc2616.txt
*/
var ReasonPhrase map[int]string

func init() {
	ReasonPhrase = map[int]string{
		http.StatusContinue:                     "Continue",
		http.StatusSwitchingProtocols:           "Switching Protocols",
		http.StatusContinue:                     "OK",
		http.StatusCreated:                      "Created",
		http.StatusAccepted:                     "Accepted",
		http.StatusNonAuthoritativeInfo:         "Non-Authoritative Information",
		http.StatusNoContent:                    "No Content",
		http.StatusResetContent:                 "Reset Content",
		http.StatusPartialContent:               "Partial Content",
		http.StatusMultipleChoices:              "Multiple Choices",
		http.StatusMovedPermanently:             "Moved Permanently",
		http.StatusFound:                        "Found",
		http.StatusSeeOther:                     "See Other",
		http.StatusNotModified:                  "Not Modified",
		http.StatusUseProxy:                     "Use Proxy",
		http.StatusTemporaryRedirect:            "Temporary Redirect",
		http.StatusBadRequest:                   "Bad Request",
		http.StatusUnauthorized:                 "Unauthorized",
		http.StatusPaymentRequired:              "Payment Required",
		http.StatusForbidden:                    "Forbidden",
		http.StatusNotFound:                     "Not Found",
		http.StatusMethodNotAllowed:             "Method Not Allowed",
		http.StatusNotAcceptable:                "Not Acceptable",
		http.StatusProxyAuthRequired:            "Proxy Authentication Required",
		http.StatusRequestTimeout:               "Request Time-out",
		http.StatusConflict:                     "Conflict",
		http.StatusGone:                         "Gone",
		http.StatusLengthRequired:               "Length Required",
		http.StatusPreconditionFailed:           "Precondition Failed",
		http.StatusRequestEntityTooLarge:        "Request Entity Too Large",
		http.StatusRequestURITooLong:            "Request-URI Too Large",
		http.StatusUnsupportedMediaType:         "Unsupported Media Type",
		http.StatusRequestedRangeNotSatisfiable: "Requested range not satisfiable",
		http.StatusExpectationFailed:            "Expectation Failed",
		http.StatusInternalServerError:          "Internal Server Error",
		http.StatusNotImplemented:               "Not Implemented",
		http.StatusBadGateway:                   "Bad Gateway",
		http.StatusServiceUnavailable:           "Service Unavailable",
		http.StatusGatewayTimeout:               "Gateway Time-out",
		http.StatusHTTPVersionNotSupported:      "HTTP Version not supported",
	}
}

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
phrase as recommended by RFC 2616 section 6.1.1, based on the
statusCode.

For example, Status(w, http.StatusBadRequest) is the equivalent
of String(w, "Bad Request", http.StatusBadRequest)
*/
func Simple(w http.ResponseWriter, statusCode int) {
	String(w, ReasonPhrase[statusCode], statusCode)
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
