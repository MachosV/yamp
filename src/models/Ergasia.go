package models

/*
Ergasia struct represents
the ergasia object of each person
*/
type Ergasia struct {
	ID        int    `json:"id"`
	IDPerson  int    `json:"idperson"`
	Perigrafi string `json:"perigrafi"`
	Date      string `json:"date"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Rank      string `json:"rank"`
	Monada    string `json:"monada"`
}
