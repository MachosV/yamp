package utils

import (
	"net/http"
	"strconv"
	"variables"
)

/*
func GetRoleNavbar
returns the appropriate navbar according to the role
*/
func GetRoleNavbar(r *http.Request) string {
	role := GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	switch roleint {
	case variables.GMT:
		return "templates/userviews/navbargmt.html"

	case variables.VICE_CAPTAIN:
		return "templates/userviews/navbarvc.html"

	case variables.CAPTAIN:
		return "templates/userviews/navbarcpt.html"

	case variables.P1G:
		return "templates/userviews/navbarp1g.html"

	case variables.YDITIS:
		return "templates/userviews/navbaryditis.html"

	case variables.DITIS:
		return "templates/userviews/navbarditis.html"

	case variables.ADMIN:
		return "templates/adminviews/navbar.html"
	}
	return "templates/userviews/navbar.html"
}
