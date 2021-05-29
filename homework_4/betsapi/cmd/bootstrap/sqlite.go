package bootstrap

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/cmd/config"
)

func Sqlite() *sql.DB {
	db, err := sql.Open("sqlite3", config.Cfg.SqliteDatabase)
	if err != nil {
		panic(err)
	}

	return db
}
