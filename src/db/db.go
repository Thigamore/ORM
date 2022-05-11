package db

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
	tables map[string]*Table
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

func (db *Db) CreateTable(tbName string)

//Adds a table that is being used to the db
func (db *Db) AddTable(table *Table) {
	db.tables[table.Name] = table
}

//Gets a table that is being used from the db
func (db *Db) GetTable(tbName string) *Table {
	tb, exists := db.tables[tbName]
	if exists {
		return tb
	} else {
		tb = &Table{
			Name: tbName,
			db:   db,
		}
		db.AddTable(tb)
		return tb
	}
}

func (db *Db) Insert(tableName string, primaryCols *[]string, cols *[]string, colTypes *[]string, newColValues *[]any) error {
	dbQuery := "INSERT INTO (?) WHERE"
}
