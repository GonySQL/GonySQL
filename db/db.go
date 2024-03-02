package database

import (
	colums_c "GonySQL/colums"
	t_table "GonySQL/table"
	"GonySQL/utils"
	"reflect"
)

type DB struct {
	Name   string
	Tables []t_table.Table
}
type DataBase interface {
	Create_table(string, []string, []reflect.Type)
}

func (db *DB) Create_table(name string, colums_name []string, colums_type []string) {
	len_n := len(colums_name)
	len_t := len(colums_type)
	Types := utils.ToType(colums_type)
	if len_n != len_t {
		panic("length of colums_name must be the same with colums_type")
	}
	var table t_table.Table
	table.Name = name
	for i := 0; i < len_n; i++ {
		var colum colums_c.Colum
		colum.Name = colums_name[i]
		colum.Type = Types[i]
		table.Colums = append(table.Colums, colum)
	}
	db.Tables = append(db.Tables, table)
}
