package views

import (
	"fmt"
	"middleware"
	"models"
	"net/http"
	"strconv"
	"utils"
)

func init() {
	GetMux().HandleFunc("/listes", middleware.WithMiddleware(listes,
		middleware.Time(),
		middleware.NoCache(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listes(w http.ResponseWriter, r *http.Request) {
	list := []models.Proswpiko{}
	list = nil
	rankmap := models.CustomMap{}
	rankmap.Init()
	clabels := make(chan []models.Ierarxia)
	defer close(clabels)
	go utils.GetDimoiries(clabels)
	if r.Method == "POST" {
		label := r.PostFormValue("label")
		labelint, _ := strconv.Atoi(label)
		list = utils.GetProswpikoList(labelint)
		for _, value := range utils.GetRankList(labelint) {
			rankmap.SetKey(value)
		}
		for index := range list {
			rankmap.Set(list[index].Rank, list[index])
		}
	}
	t, err := utils.LoadTemplates("listes",
		"templates/adminviews/listes.html",
		utils.GetRoleNavbar(r),
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	datamap := make(map[string]interface{})
	datamap["rankmap"] = rankmap
	datamap["list"] = list
	data := utils.Data{}
	data.Data = datamap
	data.Context = utils.LoadContext(r)
	datamap["labels"] = <-clabels
	t.ExecuteTemplate(w, "listes", data)
}
