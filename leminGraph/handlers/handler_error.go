package handlers

import (
	"html/template"
	"net/http"

	o "lemin/leminGraph/structures"
)

// Function for the errors pages
// redirect to the error pages
func Error(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	var message string
	// redirect to the error page
	tmpl := template.Must(template.ParseFiles(o.PathTemplate + "error.html"))
	switch status {
	case http.StatusNotFound:
		message = "Page Not Found"
	case http.StatusBadRequest:
		message = "Bad Request"
	case http.StatusInternalServerError:
		message = "Internal Server Error"
	}
	vError := *o.NewDataError(status, message)
	_ = tmpl.Execute(w, vError)
}
