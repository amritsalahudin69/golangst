package repository

import "database/sql"

type cmpusInteractor struct {
	conn *sql.DB
}
