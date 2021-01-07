package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mlshvsk/go-task-manager/config"
)

type SqlDB struct {
	Conn *sql.DB
}

func Load(dbCfg config.SqlConfig) (*SqlDB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DatabaseName,
		dbCfg.Encoding)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &SqlDB{db}, nil
}