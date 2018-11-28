package models

/*
Dynamologio struct
holds the final information
to be parsed into json
for this object
*/
type Dynamologio struct {
	Proswpiko []Proswpiko `json:"proswpiko"`
	Metaboles []AdeiaDyn  `json:"metaboles"`
	Aitiseis  []Aitisi    `json:"aitiseis"`
	Ypiresies []Ypiresia  `json:"ypiresies"`
	Anafores  []Anafora   `json:"anafores"`
	Ergasies  []Ergasia   `json:"ergasies"`
}
