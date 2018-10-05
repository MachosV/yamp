package utils

import (
	"datastorage"
	"net/http"
	"strconv"
	"variables"
)

func CanActOnPerson(r *http.Request, id int) bool {
	role := GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	if roleint >= variables.ADMIN {
		return true
	}
	labeltemp := GetSessionValue(r, "label")
	labelredis, _ := strconv.Atoi(labeltemp)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT proswpiko.id FROM proswpiko join ierarxia on proswpiko.label=ierarxia.id || proswpiko.label=ierarxia.parentid where ierarxia.id=? || ierarxia.parentid=?", labelredis, labelredis)
	var temp int
	for res.Next() {
		_ = res.Scan(&temp)
		if id == temp {
			return true
		}
	}
	return false
}