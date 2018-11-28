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
	views.GetMux().HandleFunc("/proswpiko", middleware.WithMiddleware(proswpiko,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsP1g(),
	))
}

func proswpiko(w http.ResponseWriter, r *http.Request) {
	cproswpiko := make(chan []models.Proswpiko)
	cranks := make(chan []models.Rank)
	cierarxia := make(chan []models.Ierarxia)
	defer close(cproswpiko)
	defer close(cranks)
	defer close(cierarxia)
	go getProswpiko(cproswpiko)
	go getRanks(cranks)
	go utils.GetDimoiries(cierarxia)
	proswpikolist := <-cproswpiko
	ranks := <-cranks
	ierarxia := <-cierarxia
	datamap := make(map[string]interface{})
	datamap["proswpiko"] = proswpikolist
	datamap["ranks"] = ranks
	datamap["ierarxia"] = ierarxia
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = datamap
	t, err := utils.LoadTemplates("proswpiko",
		"templates/adminviews/proswpiko.html",
		utils.GetRoleNavbar(r),
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "proswpiko", data)
}

func getProswpiko(c chan []models.Proswpiko) {
	proswpikoArray := []models.Proswpiko{}
	proswpiko := models.Proswpiko{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT id,pname,surname,rank from proswpiko_sorted;")
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Name,
			&proswpiko.Surname,
			&proswpiko.Rank,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	res.Close()
	c <- proswpikoArray
}

func getRanks(c chan []models.Rank) {
	ranksArray := []models.Rank{}
	rank := models.Rank{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT * from ranks;")
	for res.Next() {
		_ = res.Scan(
			&rank.ID,
			&rank.Rank,
		)
		ranksArray = append(ranksArray, rank)
	}
	res.Close()
	c <- ranksArray
}
