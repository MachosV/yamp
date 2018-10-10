package adminviews

import (
	"datastorage"
	"encoding/json"
	"fmt"
	"log"
	"middleware"
	"models"
	"net/http"
	"strconv"
	"time"
	"utils"
	"variables"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/getdyn", middleware.WithMiddleware(getdyn,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func getdyn(w http.ResponseWriter, r *http.Request) {
	label := r.URL.Query().Get("label")
	labelform, _ := strconv.Atoi(label)
	date := r.URL.Query().Get("date")
	var data string
	data = "Not found"
	role := utils.GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	if label == "all" && roleint == variables.ADMIN {
		data = getfulldyn(date)
	} else if label != "all" && roleint == variables.ADMIN {
		data = getdynlabel(date, labelform)
	} else {
		if utils.CheckLabelAuthed(r, labelform) {
			data = getdynlabel(date, labelform)
		} else {
			data = "Unauthenticated"
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", data)
}

func getfulldyn(d string) string {
	var dynamologio models.Dynamologio
	c := make(chan []models.AdeiaDyn)
	cmin := make(chan []models.MinDynRecord)
	cminmetaboles := make(chan []models.MinDynAdeiaRecord)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	dtemp, _ := time.Parse("02/01/2006", d)
	datefordb := fmt.Sprintf("%d/%d/%d", dtemp.Year(), dtemp.Month(), dtemp.Day())
	go utils.GetDynAdeies(datefordb, c)
	go utils.GetDynMinAll(datefordb, cmin)
	go utils.GetDynMinAdeiesAll(datefordb, cminmetaboles)
	res, err := dbc.Query("select proswpiko_sorted.id,pname,surname,rank,perigrafi from proswpiko_sorted where proswpiko_sorted.id not in (select idperson from adeies where ? between adeies.start and adeies.end);", datefordb)
	if err != nil {
		log.Println(err)
		return "Σφάλμα ανάκτησης δυναμολογίου"
	}
	var proswpiko models.Proswpiko
	var proswpikoArray []models.Proswpiko
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Name,
			&proswpiko.Surname,
			&proswpiko.Rank,
			&proswpiko.Label,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	res.Close()
	res, err = dbc.Query("select distinct ranks.id,ranks.rank from proswpiko_sorted join ranks on ranks.rank = proswpiko_sorted.rank where proswpiko_sorted.id not in (select idperson from adeies where ? between adeies.start and adeies.end) order by ranks.id DESC;", d)
	if err != nil {
		log.Println(err)
		return "Σφάλμα ανάκτησης βαθμών δυναμολογίου"
	}
	var rank models.Rank
	var rankmap models.CustomMap
	rankmap.Init()
	for res.Next() {
		_ = res.Scan(
			&rank.ID,
			&rank.Rank,
		)
		rankmap.SetKey(rank.Rank)
	}
	res.Close()
	for index := range proswpikoArray {
		rankmap.Set(proswpikoArray[index].Rank, proswpikoArray[index])
	}
	dynamologio.Rankmap = rankmap
	dynamologio.Metaboles = <-c
	dynamologio.MetabolesMin = <-cminmetaboles
	jsonString, err := json.MarshalIndent(dynamologio, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(jsonString[:])
}

func getdynlabel(d string, label int) string {
	var dynamologio models.Dynamologio
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	c := make(chan []models.AdeiaDyn)
	cmin := make(chan []models.MinDynRecord)
	cminmetaboles := make(chan []models.MinDynAdeiaRecord)
	dtemp, _ := time.Parse("02/01/2006", d)
	datefordb := fmt.Sprintf("%d/%d/%d", dtemp.Year(), dtemp.Month(), dtemp.Day())
	go utils.GetDynAdeiesLabeled(datefordb, label, c)
	go utils.GetDynMinLabel(datefordb, label, cmin)
	go utils.GetDynMinAdeiesLabel(datefordb, label, cminmetaboles)
	res, err := dbc.Query("select proswpiko_sorted.id,pname,surname,rank,ierarxia.perigrafi from proswpiko_sorted join ierarxia on ierarxia.perigrafi = proswpiko_sorted.perigrafi where proswpiko_sorted.id not in (select idperson from adeies where ? between adeies.start and adeies.end) and (ierarxia.id = ? ||ierarxia.parentid = ?);", datefordb, label, label)
	if err != nil {
		log.Println(err)
		return "Σφάλμα ανάκτησης δυναμολογίου"
	}
	var proswpiko models.Proswpiko
	var proswpikoArray []models.Proswpiko
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Name,
			&proswpiko.Surname,
			&proswpiko.Rank,
			&proswpiko.Label,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	res.Close()
	res, err = dbc.Query("select distinct ranks.id,ranks.rank from proswpiko_sorted join ranks on ranks.rank = proswpiko_sorted.rank where proswpiko_sorted.id not in (select idperson from adeies where ? between adeies.start and adeies.end) order by ranks.id DESC;", d)
	if err != nil {
		log.Println(err)
		return "Σφάλμα ανάκτησης βαθμών δυναμολογίου"
	}
	var rank models.Rank
	var rankmap models.CustomMap
	rankmap.Init()
	for res.Next() {
		_ = res.Scan(
			&rank.ID,
			&rank.Rank,
		)
		rankmap.SetKey(rank.Rank)
	}
	res.Close()
	for index := range proswpikoArray {
		rankmap.Set(proswpikoArray[index].Rank, proswpikoArray[index])
	}
	dynamologio.Rankmap = rankmap
	dynamologio.Metaboles = <-c
	dynamologio.MetabolesMin = <-cminmetaboles
	jsonString, err := json.MarshalIndent(dynamologio, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(jsonString[:])
}
