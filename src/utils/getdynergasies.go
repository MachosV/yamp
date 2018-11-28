package utils

import (
	"bytes"
	"datastorage"
	"log"
	"models"
)

/*
GetDynErgasiesAll function
fetches all the different objects in the database
of the table anafores for a certain date
*/
func GetDynErgasiesAll(d string, c chan []models.Ergasia) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT ergasies.perigrafi,ergasies.date,proswpiko.id,proswpiko.name,")
	buffer.WriteString("proswpiko.surname, ranks.rank, ierarxia.perigrafi as monada ")
	buffer.WriteString("FROM ergasies ")
	buffer.WriteString("JOIN proswpiko ON ergasies.idperson = proswpiko.id ")
	buffer.WriteString("JOIN ranks ON proswpiko.rank = ranks.id ")
	buffer.WriteString("JOIN ierarxia ON proswpiko.label = ierarxia.id ")
	buffer.WriteString("WHERE ergasies.date = ? ORDER BY ranks.id DESC")
	res, err := dbc.Query(buffer.String(), d)
	if err != nil {
		log.Println(err)
	}
	var ergasia models.Ergasia
	var ergasies []models.Ergasia
	for res.Next() {
		_ = res.Scan(
			&ergasia.Perigrafi,
			&ergasia.Date,
			&ergasia.IDPerson,
			&ergasia.Name,
			&ergasia.Surname,
			&ergasia.Rank,
			&ergasia.Monada,
		)
		ergasia.Date = models.DateBuilder(ergasia.Date)
		ergasies = append(ergasies, ergasia)
	}
	res.Close()
	c <- ergasies
}

/*
GetDynErgasiesLabel function
returns all anafores objects for a certain label
*/
func GetDynErgasiesLabel(d string, label int, c chan []models.Ergasia) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT ergasies.perigrafi,ergasies.date,proswpiko.id,proswpiko.name,")
	buffer.WriteString("proswpiko.surname, ranks.rank, ierarxia.perigrafi as monada ")
	buffer.WriteString("FROM ergasies ")
	buffer.WriteString("JOIN proswpiko ON ergasies.idperson = proswpiko.id ")
	buffer.WriteString("JOIN ranks ON proswpiko.rank = ranks.id ")
	buffer.WriteString("JOIN ierarxia ON proswpiko.label = ierarxia.id ")
	buffer.WriteString("WHERE ergasies.date = ? and (ierarxia.id =? || ierarxia.parentid = ?) ORDER BY ranks.id DESC")
	res, _ := dbc.Query(buffer.String(), d, label, label)
	var ergasia models.Ergasia
	var ergasies []models.Ergasia
	for res.Next() {
		_ = res.Scan(
			&ergasia.Perigrafi,
			&ergasia.Date,
			&ergasia.IDPerson,
			&ergasia.Name,
			&ergasia.Surname,
			&ergasia.Rank,
			&ergasia.Monada,
		)
		ergasia.Date = models.DateBuilder(ergasia.Date)
		ergasies = append(ergasies, ergasia)
	}
	res.Close()
	c <- ergasies
}
