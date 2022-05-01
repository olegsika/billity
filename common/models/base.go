package models

import "github.com/go-pg/pg"

type Base interface {
	isUpdate() bool
	setTimeStamp()
	updateTime()
}

func Insert(model Base, db *pg.DB) error {
	model.setTimeStamp()

	return db.Insert(model)
}

func Save(model Base, db *pg.DB) error {
	if model.isUpdate() {
		model.updateTime()

		var _, err = db.Model(model).WherePK().Update()
		return err
	}

	return Insert(model, db)
}
