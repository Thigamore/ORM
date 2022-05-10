package src

import (
	"reflect"
)

type Table struct {
	Name string
	db   *Db
}

func InitTable(tableName string, db *Db) *Table {
	table := &Table{Name: tableName, db: db}
	return table
}

//Saves a particular object to a db
func (tb *Table) Save(obj any) error {
	//Gets information about the object
	objType := reflect.TypeOf(obj)
	fieldNum := objType.NumField()
	fieldCols := make([]string, fieldNum)
	colValues := make([]any, fieldNum)
	colTypes := make([]string, fieldNum)
	for i := 0; i < fieldNum; i++ {
		fieldCols[i] = objType.Field(i).Tag.Get("orm")
		colValues[i] = reflect.ValueOf(obj).FieldByName(fieldCols[i])
		colTypes[i] = objType.Field(i).Type.Name()
	}

	return tb.db.Insert(tb.Name, &fieldCols, &colTypes, &colValues)
}
