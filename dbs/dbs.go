package databases

import database "GonySQL/db"

func (dbs *DBS) Create_db(name string) database.DB {
	var db database.DB
	db.Name = name
	dbs.databases = append(dbs.databases, db)
	return db
}

type DBS struct {
	databases []database.DB
}
type GonySQL interface {
	Create_db(string)
}
