package src

import (
	"database/sql"
	"log"
	"strings"

	"github.com/Thigamore/ORM/src"

	_ "github.com/go-sql-driver/mysql"
)

//Db Wrapper
type Db struct {
	db     *sql.DB
	tables []*Table
}

func InitDb(dbName string, conf src.Config) (*Db, error) {
	log.Println("Initializing Db")
	db := &Db{}

	dbSql, err := sql.Open("mysql", conf.Name+":"+conf.Password+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		return nil, err
	}

	db.db = dbSql

	return db, nil
}

func (db *Db) TableExists(tableName string) (bool, error) {
	dbQuery := "SHOW TABLES LIKE '" + tableName + "';"
	rs, err := db.db.Query(dbQuery)
	var result string
	if err != nil {
		return false, err
	}
	for rs.Next() {
		rs.Scan(&result)
		if strings.ToLower(tableName) == result {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}

func (db *Db) Insert(tableName string, primaryCols *[]string, cols *[]string, colTypes *[]string, newColValues *[]any) error {
	dbQuery := "INSERT INTO (?) WHERE"
}
