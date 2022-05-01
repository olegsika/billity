package db

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"time"

	// DB adapter
	_ "github.com/lib/pq"
)

type DB interface {
	orm.DB
}

// NewGoPG New creates new database connection to a postgres database
// Function panics if it can't connect to database
func NewGoPG(psn string) (*pg.DB, error) {
	u, err := pg.ParseURL(psn)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(u).WithTimeout(time.Second * 50)

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
