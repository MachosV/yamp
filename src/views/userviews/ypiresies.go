package userviews

import (
	"fmt"
	"middleware"
	"models"
	"net/http"
	"strconv"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/ypiresies", middleware.WithMiddleware(ypiresies,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func ypiresies(w http.ResponseWriter, r *http.Request) {
	clabels := make(chan []models.Ierarxia)
	ctypoiypiresiwn := make(chan []models.TyposYpiresias)
	defer close(clabels)
	defer close(ctypoiypiresiwn)
	labelredis, _ := strconv.Atoi(utils.GetSessionValue(r, "label"))
	go utils.GetLabels(labelredis, clabels)
	datamap := make(map[string]interface{})
	if r.Method == "POST" {
		label := r.PostFormValue("label")
		labelform, _ := strconv.Atoi(label)
		if utils.CheckLabelAuthed(r, labelform) {
			go utils.GetTypoiYpiresiwn(labelform, ctypoiypiresiwn)
			datamap["ypiresies"] = <-ctypoiypiresiwn
		}
	}
	t, err := utils.LoadTemplates("ypiresies",
		"templates/userviews/ypiresies.html",
		utils.GetRoleNavbar(r),
		"templates/userviews/header.html",
		"templates/userviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	datamap["labels"] = <-clabels
	data := utils.Data{}
	data.Data = datamap
	data.Context = utils.LoadContext(r)
	t.ExecuteTemplate(w, "ypiresies", data)
}
