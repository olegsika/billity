package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

type DBUsers struct {
}

func NewUsersDB() *DBUsers {
	return &DBUsers{}
}

func (u *DBUsers) CreateUser(user *models.User, db *pg.DB) error {
	return models.Save(user, db)
}

func (u *DBUsers) UpdateUser(user *models.User, db *pg.DB) error {
	return models.Save(user, db)
}

func (u *DBUsers) GetUser(msisdn string, isForUpdate bool, db *pg.DB) (models.User, error) {
	user := models.User{}

	query := db.Model(&user).
		Where("msisdn = ?", msisdn).
		Where("deleted_at is null")

	if isForUpdate {
		query.Column("id", "created_at", "updated_at")
	}

	err := query.Select()

	if err != nil {
		return user, nil
	}

	return user, err
}
