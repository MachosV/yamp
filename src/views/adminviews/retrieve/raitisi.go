package retrieve

import (
	"fmt"
	"middleware"
	"models"
	"net/http"
	"strconv"
	"utils"
	"variables"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/raitisi", middleware.WithMiddleware(raitisi,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func raitisi(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)
	if !utils.CanActOnAitisi(r, idint) {
		http.Redirect(w, r, "/notfound", http.StatusMovedPermanently)
		return
	}
	cpersonaitisi := make(chan *models.PersonAitisi)
	cypografesaitisewn := make(chan []models.YpografiAitisis)
	defer close(cpersonaitisi)
	defer close(cypografesaitisewn)
	go utils.GetPersonAitisi(idint, cpersonaitisi)
	go utils.GetYpografesAitisis(idint, cypografesaitisewn)
	datamap := make(map[string]interface{})
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	datamap["PersonAitisi"] = <-cpersonaitisi
	datamap["Ypografes"] = <-cypografesaitisewn
	data.Data = datamap
	role := utils.GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	var navbar string
	if roleint >= variables.ADMIN {
		navbar = "./templates/adminviews/navbar.html"
	} else {
		navbar = "./templates/userviews/navbar.html"
	}
	t, err := utils.LoadTemplates("raitisi",
		"./templates/adminviews/raitisi.html",
		"./templates/adminviews/header.html",
		"./templates/adminviews/footer.html",
		navbar,
	)
	if err != nil {
		fmt.Fprintf(w, "Error -> %s", err)
		return
	}
	t.ExecuteTemplate(w, "raitisi", data)
	//fmt.Fprintf(w, "Not yet implemented")
}
