package dbrepo

import (
	"database/sql"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {

	return m.DB

}

func printRows(rows *sql.Rows) {

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Println("error no able to print rows for coffees")
	}
	for _, s := range colTypes {
		log.Println("cols type:", s.Name(), "  ", s.DatabaseTypeName(), " to gp ", s.ScanType())
	}
}


