package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

// DBCallHistory struct for DBCallHistory
type DBCallHistory struct {
}

// NewCallHistoryDB init DBCallHistory
func NewCallHistoryDB() *DBCallHistory {
	return &DBCallHistory{}
}

// CreateCallHistory the function create the call history
func (c *DBCallHistory) CreateCallHistory(history *models.CallHistory, db *pg.DB) error {
	return models.Save(history, db)
}
