# respond

Utility functions for responding to net/http requests

Feel free to see the [docs for respond](https://godoc.org/github.com/jeffbmartinez/respond).

## Usage examples

### Respond with string

```
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	respond.String(w, http.StatusOK, "This is the response string")
}
```

### Respond with json

This will automatically set `Content-Type` header as appropriate.

```
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	myObject := map[string]interface{}{
		// ...
	}

	if err := respond.JSON(w, http.StatusOK, myObject); err != nil {
		// Problem serializing object to json, handle error as appropriate
		respond.Simple(w, http.StatusInternalServerError)
		return
	}
}
```

### Respond with html

This will automatically set `Content-Type` header as appropriate.

```
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	htmlResponseString := "<html>...</html>"
	respond.HTML(w, http.StatusOK, htmlResponseString)
}
```

### Respond with html template

```
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	templateName := "index.html"
	data := /* data used to populate template */

	if err := respond.HTMLTemplate(w, http.StatusOK, templateName, data); err != nil {
	    // template rendering problem :(
	    // maybe log it and return 500 internal server error

	    respond.Simple(w, http.StatusInternalServerError)
	    return
	}
}
```
