package src

type Table struct {
	Name string
}

func InitTable(tableName string) *Table {
	return &Table{Name: tableName}
}

func (tb *Table) Save()
