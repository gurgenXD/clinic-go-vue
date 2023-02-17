package repository

import (
	"database/sql"
	"time"
)

type specialist struct {
	Uuid          string         `db:"uuid"`
	FirstName     string         `db:"first_name"`
	LastName      string         `db:"last_name"`
	MiddleName    sql.NullString `db:"middle_name"`
	StartWorkDate time.Time      `db:"start_work_date"`
	Photo         sql.NullString `db:"photo"`
	Education     string         `db:"education"`
	Activity      sql.NullString `db:"activity"`
	Description   sql.NullString `db:"description"`
	Titles        sql.NullString `db:"titles"`

	OnMain    bool `db:"on_main"`
	IsActive  bool `db:"is_active"`
	CanOnline bool `db:"can_online"`

	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
