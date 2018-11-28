package adminviews

import (
	"fmt"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/typoiadeiwn", middleware.WithMiddleware(listtypoiadeiwn,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listtypoiadeiwn(w http.ResponseWriter, r *http.Request) {
	ctypoiadeias := make(chan []models.TyposAdeias)
	ccategoryadeias := make(chan []models.CategoryAdeias)
	defer close(ctypoiadeias)
	defer close(ccategoryadeias)
	go utils.GetTypoiAdeiwn(ctypoiadeias)
	go utils.GetCategoriesAdeiwn(ccategoryadeias)
	categories := <-ccategoryadeias
	typoiadeiwn := <-ctypoiadeias
	adeiesmap := models.CustomMap{}
	adeiesmap.Init()
	for _, value := range categories {
		adeiesmap.SetKey(value.Category)
	}
	for index := range typoiadeiwn {
		adeiesmap.Set(typoiadeiwn[index].Category, typoiadeiwn[index])
	}
	datamap := make(map[string]interface{})
	datamap["adeies"] = adeiesmap
	datamap["categories"] = categories
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = datamap
	t, err := utils.LoadTemplates("typoiadeiwn",
		"templates/adminviews/typoiadeiwn.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html",
		utils.GetRoleNavbar(r))
	if err != nil {
		fmt.Fprintf(w, "Error->%s", err)
		return
	}
	t.ExecuteTemplate(w, "typoiadeiwn", data)
}
