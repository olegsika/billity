package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

type DBCallHistory struct {
}

func NewCallHistoryDB() *DBCallHistory {
	return &DBCallHistory{}
}

func (c *DBCallHistory) CreateCallHistory(history *models.CallHistory, db *pg.DB) error {
	return models.Save(history, db)
}
