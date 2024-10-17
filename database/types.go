package database

import "database/sql"

type Backend struct {
	DB *sql.DB
}
