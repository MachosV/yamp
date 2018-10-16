package userviews

import (
	"fmt"
	"middleware"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/dashboard", middleware.WithMiddleware(Dashboard,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

/*
Dashboard function of package userviews
provides the main page of the user when logged in
*/
func Dashboard(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	data := utils.Data{}
	data.Context = context
	t, err := utils.LoadTemplates("dashboard",
		"templates/userviews/dashboard.html",
		"templates/userviews/navbar.html",
		"templates/userviews/header.html",
		"templates/userviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "dashboard", data)

}
