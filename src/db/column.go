package db

import (
	colArgs "github.com/Thigamore/ORM/src/db/args/columnArgs"
)

type column struct {
	name     string
	dataType string
	args     []colArgs.ColumnArg
}

//The arguments for the column
const ()

func CreateColumn(name string, dataType string) *column {
	return &column{
		name:     name,
		dataType: dataType,
	}
}
