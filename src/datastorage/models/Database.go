package models

/*
DatabaseConf object that gets loaded from databases.conf
is used by DataRouter object to create database connections
*/
type DatabaseConf struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Link     string `json:"link"`
	Database string `json:"database"`
	Type     string `json:"type"`
}
