package adminviews

import (
	"bytes"
	"datastorage"
	"fmt"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/listaitiseis", middleware.WithMiddleware(listaitiseis,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.NoCache(),
		middleware.IsCaptain(),
	))
}

func listaitiseis(w http.ResponseWriter, r *http.Request) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	ccountaitiseis := make(chan int)
	defer close(ccountaitiseis)
	go utils.CountAitiseis(ccountaitiseis, utils.GetSessionValue(r, "role"))
	buffer.WriteString("SELECT aitiseis.id, idperson, aitiseis.perigrafi, date,")
	buffer.WriteString("name, surname, ranks.rank, ierarxia.perigrafi FROM aitiseis ")
	buffer.WriteString("JOIN proswpiko on idperson = proswpiko.id ")
	buffer.WriteString("JOIN ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("JOIN ierarxia on proswpiko.label = ierarxia.id ")
	buffer.WriteString("WHERE aitiseis.id not in (SELECT idaitisi from ypografes_aitisewn where signedas =?) ")
	buffer.WriteString("ORDER BY aitiseis.id ASC, proswpiko.rank DESC, date ASC LIMIT 4;")
	res, err := dbc.Query(buffer.String(), utils.GetSessionValue(r, "role"))
	if err != nil {
		utils.RedirectWithError(w, r, utils.RedirectByRole(r), "Σφάλμα ανάκτησης αιτήσεων", err)
		return
	}
	var aitisi models.Aitisi
	var aitiseis []models.Aitisi
	for res.Next() {
		_ = res.Scan(
			&aitisi.ID,
			&aitisi.IDPerson,
			&aitisi.Perigrafi,
			&aitisi.Date,
			&aitisi.Name,
			&aitisi.Surname,
			&aitisi.Rank,
			&aitisi.Monada,
		)
		aitisi.Date = models.DateBuilder(aitisi.Date)
		aitiseis = append(aitiseis, aitisi)
	}
	datamap := make(map[string]interface{})
	datamap["Context"] = utils.LoadContext(r)
	datamap["Aitiseis"] = aitiseis
	datamap["CountAitiseis"] = <-ccountaitiseis
	t, err := utils.LoadTemplates("listaitiseis",
		"templates/adminviews/listaitiseis.html",
		utils.GetRoleNavbar(r),
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "listaitiseis", datamap)
}
