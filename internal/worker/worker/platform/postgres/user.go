package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

// DBUser struct for DBUser
type DBUser struct {
}

// NewUserDB init DBUser
func NewUserDB() *DBUser {
	return &DBUser{}
}

// GetUser the function get user by msisdn
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

// UpdateBalance the function update balance for user
func (u *DBUser) UpdateBalance(user *models.User, db *pg.DB) error {
	return models.Save(user, db)
}
