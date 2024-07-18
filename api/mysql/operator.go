package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sreok/kube-go/configs"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", configs.Db)
	if err != nil {
		panic(err)
	}
}

type ApiGroups struct {
	Id         int
	Name       string
	ShortNames string
	ApiVersion string
	Namespaced bool
	Kind       string
}

func InsertApiGroups() {
	sql := "insert into " + configs.DbTable + " (name, ShortNames, ApiVersion, Namespaced, Kind) values (?, ?, ?, true, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
}
