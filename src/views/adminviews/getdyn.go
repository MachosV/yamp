package adminviews

import (
	"encoding/json"
	"fmt"
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
		middleware.NoCache(),
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
	if label == "all" && roleint >= variables.P1G {
		data = getfulldyn(date)
	} else if label != "all" && roleint >= variables.P1G {
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
	cproswpiko := make(chan []models.Proswpiko)
	cadeies := make(chan []models.AdeiaDyn)
	cypiresies := make(chan []models.Ypiresia)
	caitiseis := make(chan []models.Aitisi)
	canafores := make(chan []models.Anafora)
	cergasies := make(chan []models.Ergasia)
	defer close(cproswpiko)
	defer close(cadeies)
	defer close(cypiresies)
	defer close(caitiseis)
	defer close(canafores)
	dtemp, _ := time.Parse("02/01/2006", d)
	datefordb := fmt.Sprintf("%d/%d/%d", dtemp.Year(), dtemp.Month(), dtemp.Day())
	go utils.GetDynProswpikoAll(datefordb, cproswpiko)
	go utils.GetDynAdeies(datefordb, cadeies)
	go utils.GetDynYpiresiesAll(datefordb, cypiresies)
	go utils.GetDynAitiseisAll(datefordb, caitiseis)
	go utils.GetDynAnaforesAll(datefordb, canafores)
	go utils.GetDynErgasiesAll(datefordb, cergasies)
	dynamologio.Proswpiko = <-cproswpiko
	dynamologio.Metaboles = <-cadeies
	dynamologio.Ypiresies = <-cypiresies
	dynamologio.Aitiseis = <-caitiseis
	dynamologio.Anafores = <-canafores
	dynamologio.Ergasies = <-cergasies
	jsonString, err := json.MarshalIndent(dynamologio, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(jsonString[:])
}

func getdynlabel(d string, label int) string {
	var dynamologio models.Dynamologio
	cproswpiko := make(chan []models.Proswpiko)
	cadeies := make(chan []models.AdeiaDyn)
	cypiresies := make(chan []models.Ypiresia)
	caitiseis := make(chan []models.Aitisi)
	canafores := make(chan []models.Anafora)
	cergasies := make(chan []models.Ergasia)
	defer close(cproswpiko)
	defer close(cadeies)
	defer close(cypiresies)
	defer close(caitiseis)
	defer close(canafores)
	defer close(cergasies)
	dtemp, _ := time.Parse("02/01/2006", d)
	datefordb := fmt.Sprintf("%d/%d/%d", dtemp.Year(), dtemp.Month(), dtemp.Day())
	go utils.GetDynProswpikoLabel(datefordb, label, cproswpiko)
	go utils.GetDynAdeiesLabeled(datefordb, label, cadeies)
	go utils.GetDynYpiresiesLabel(datefordb, label, cypiresies)
	go utils.GetDynAitiseisLabel(datefordb, label, caitiseis)
	go utils.GetDynAnaforesLabel(datefordb, label, canafores)
	go utils.GetDynErgasiesLabel(datefordb, label, cergasies)
	dynamologio.Proswpiko = <-cproswpiko
	dynamologio.Metaboles = <-cadeies
	dynamologio.Ypiresies = <-cypiresies
	dynamologio.Aitiseis = <-caitiseis
	dynamologio.Anafores = <-canafores
	dynamologio.Ergasies = <-cergasies
	jsonString, err := json.MarshalIndent(dynamologio, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(jsonString[:])
}
