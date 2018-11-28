package utils

import (
	"net/http"
	"strconv"
	"variables"
)

/*
RedirectByRole is a helper function
that returns a url string according to the role of the caller
*/
func RedirectByRole(r *http.Request) string {
	role := GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	isAuthenticated := GetSessionValue(r, "isAuthenticated")
	if isAuthenticated == "false" || role == "" {
		return "/"
	}
	if roleint >= variables.ADMIN {
		return "/diaxeiristiko"
	}
	return "/dashboard"
}
