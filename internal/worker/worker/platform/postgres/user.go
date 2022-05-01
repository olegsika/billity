package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

type DBUser struct {
}

func NewUserDB() *DBUser {
	return &DBUser{}
}

func (u *DBUser) GetUser(msisdn string, db *pg.DB) (models.User, error) {
	user := models.User{}

	err := db.Model(&user).
		Where("msisdn = ?", msisdn).
		Where("deleted_at is null").
		Select()

	if err != nil {
		return user, nil
	}

	return user, err
}

func (u *DBUser) UpdateBalance(user *models.User, db *pg.DB) error {
	return models.Save(user, db)
}
