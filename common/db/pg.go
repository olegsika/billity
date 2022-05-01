package db

import (
	"billity/common/utils"
	"database/sql"
	"github.com/cenkalti/backoff"
	"time"
)

const timeoutDuration = "2s"

func InitDBClient(psn string) (db *sql.DB) {
	timeout, err := time.ParseDuration(timeoutDuration)
	utils.CheckErr(err)

	operation := func() error {
		db, err = sql.Open("postgres", psn)
		if err != nil {
			return err
		}
		if err := db.Ping(); err != nil {
			return err
		}
		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = timeout

	err = backoff.Retry(operation, b)
	utils.CheckErr(err)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
