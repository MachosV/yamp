package adminviews

import (
	"fmt"
	"middleware"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/diaxeiristiko", middleware.WithMiddleware(Diaxeiristiko,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

/*
Diaxeiristiko function
serves as the function to the url /dashboardadmin
*/
func Diaxeiristiko(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	data := utils.Data{}
	data.Context = context
	t, err := utils.LoadTemplates("dashboard",
		"./templates/adminviews/dashboard.html",
		utils.GetRoleNavbar(r),
		"./templates/adminviews/header.html",
		"./templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Error->%s", err)
	}
	t.ExecuteTemplate(w, "dashboard", data)
}
