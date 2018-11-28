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
	views.GetMux().HandleFunc("/stelexh", middleware.WithMiddleware(stelexh,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func stelexh(w http.ResponseWriter, r *http.Request) {
	clabels := make(chan []models.Ierarxia)
	defer close(clabels)
	list := []models.Proswpiko{}
	list = nil
	rankmap := models.CustomMap{}
	rankmap.Init()
	labelredis, _ := strconv.Atoi(utils.GetSessionValue(r, "label"))
	go utils.GetLabels(labelredis, clabels)
	if r.Method == "POST" {
		label := r.PostFormValue("label")
		labelform, _ := strconv.Atoi(label)
		if utils.CheckLabelAuthed(r, labelform) {
			list = utils.GetProswpikoList(labelform)
			for _, value := range utils.GetRankList(labelform) {
				rankmap.SetKey(value)
			}
			for index := range list {
				rankmap.Set(list[index].Rank, list[index])
			}
		}
	}
	t, err := utils.LoadTemplates("stelexh",
		"templates/userviews/stelexh.html",
		utils.GetRoleNavbar(r),
		"templates/userviews/header.html",
		"templates/userviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	datamap := make(map[string]interface{})
	datamap["rankmap"] = rankmap
	datamap["list"] = list
	datamap["labels"] = <-clabels
	data := utils.Data{}
	data.Data = datamap
	data.Context = utils.LoadContext(r)
	t.ExecuteTemplate(w, "stelexh", data)
}
