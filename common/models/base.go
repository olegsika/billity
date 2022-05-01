package models

import "github.com/go-pg/pg"

// Base the interface for all the models
type Base interface {
	isUpdate() bool
	setTimeStamp()
	updateTime()
}

// Insert insert model to DB
func Insert(model Base, db *pg.DB) error {
	model.setTimeStamp()

	return db.Insert(model)
}

// Save insert or update model to DB
func Save(model Base, db *pg.DB) error {
	if model.isUpdate() {
		model.updateTime()

		var _, err = db.Model(model).WherePK().Update()
		return err
	}

	return Insert(model, db)
}
