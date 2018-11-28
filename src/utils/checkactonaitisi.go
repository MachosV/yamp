package utils

import (
	"bytes"
	"datastorage"
	"net/http"
	"strconv"
	"variables"
)

/*
CanActOnAitisi function
is used to check whether a request can manipulate
or view the requested aitisi object
*/
func CanActOnAitisi(r *http.Request, id int) bool {
	role := GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	if roleint >= variables.P1G {
		return true
	}
	labeltemp := GetSessionValue(r, "label")
	labelredis, _ := strconv.Atoi(labeltemp)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT aitiseis.id from aitiseis ")
	buffer.WriteString("JOIN proswpiko on aitiseis.idperson = proswpiko.id ")
	buffer.WriteString("JOIN ierarxia on ")
	buffer.WriteString("proswpiko.label = ierarxia.id || ")
	buffer.WriteString("proswpiko.label = ierarxia.parentid ")
	buffer.WriteString("where ierarxia.id = ? || ierarxia.parentid = ?;")
	res, _ := dbc.Query(buffer.String(), labelredis, labelredis)
	defer res.Close()
	var temp int
	for res.Next() {
		_ = res.Scan(&temp)
		if id == temp {
			return true
		}
	}
	return false
}
