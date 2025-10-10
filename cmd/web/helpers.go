package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError prints the error and stack trace to the log and sends a 500 Internal Server Error response.
func (app *application) serverError(w http.ResponseWriter, err error) {

	// Print the error and stack trace to the log
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a response with the given status code and text.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound sends a 404 Not Found response.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
