package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

type DBUsage struct {
}

func NewUsageDB() *DBUsage {
	return &DBUsage{}
}

func (u *DBUsage) GetBalance(msisdn string, db *pg.DB) (float64, error) {
	var balance float64

	err := db.Model((*models.User)(nil)).
		Where("msisdn = ?", msisdn).
		Column("balance").
		Select(&balance)

	if err != nil {
		return 0, err
	}

	return balance, nil
}
