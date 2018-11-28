package views

import (
	"fmt"
	"middleware"
	"net/http"
)

func init() {
	GetMux().HandleFunc("/csrfdenied", middleware.WithMiddleware(csrfdenied,
		middleware.Time(),
	))
}

func csrfdenied(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Csrf token invalid")
}
