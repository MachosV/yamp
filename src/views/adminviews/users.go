package adminviews

import (
	"datastorage"
	"fmt"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/listusers", middleware.WithMiddleware(listusers,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listusers(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	mysqlclient := db.GetMysqlClient()
	//var buffer bytes.Buffer
	//buffer.WriteString("SELECT accounts.id,username,rolestring")
	res, err := mysqlclient.Query("SELECT id,username,rolestring from accounts join roles on accounts.role=roles.role order by roles.role desc, accounts.username asc")
	defer res.Close()
	if err != nil {
		context.Message = err.Error()
		http.Redirect(w, r, "/listusers", http.StatusMovedPermanently)
	}
	var users []models.User
	var user models.User
	for res.Next() {
		err = res.Scan(
			&user.ID,
			&user.Username,
			&user.RealRole)
		users = append(users, user)
	}
	res, err = mysqlclient.Query("SELECT role,rolestring from roles ORDER BY role ASC")
	var role models.Role
	var roles []models.Role
	for res.Next() {
		_ = res.Scan(
			&role.Role,
			&role.Rolestring,
		)
		roles = append(roles, role)
	}
	data := utils.Data{}
	datamap := make(map[string]interface{})
	data.Context = context
	datamap["Users"] = users
	datamap["Roles"] = roles
	data.Data = datamap
	t, err := utils.LoadTemplates("users",
		"templates/adminviews/users.html",
		utils.GetRoleNavbar(r),
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "users", data)
}
