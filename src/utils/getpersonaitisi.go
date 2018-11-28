package utils

import (
	"bytes"
	"datastorage"
	"models"
)

/*
GetPersonAitisi function
returns asynchronously the personaitisi object requested
*/
func GetPersonAitisi(id int, c chan *models.PersonAitisi) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT aitiseis.id,aitiseis.perigrafi,aitiseis.date,ranks.rank,ierarxia.perigrafi,name,surname ")
	buffer.WriteString("FROM aitiseis ")
	buffer.WriteString("JOIN proswpiko on aitiseis.idperson = proswpiko.id ")
	buffer.WriteString("JOIN ierarxia on proswpiko.label = ierarxia.id as monada ")
	buffer.WriteString("JOIN ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("WHERE aitiseis.id = ?")
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query(buffer.String(), id)
	defer res.Close()
	if err != nil {
		c <- nil
		return
	}
	var personaitisi models.PersonAitisi
	if res.Next() {
		_ = res.Scan(
			&personaitisi.ID,
			&personaitisi.Perigrafi,
			&personaitisi.Date,
			&personaitisi.Rank,
			&personaitisi.Monada,
			&personaitisi.Name,
			&personaitisi.Surname,
		)
	}
	c <- &personaitisi
}
