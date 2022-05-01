package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

// DBUsers struct for DBUsers
type DBUsers struct {
}

// NewUsersDB init DBUsers
func NewUsersDB() *DBUsers {
	return &DBUsers{}
}

// CreateUser the function create the User
func (u *DBUsers) CreateUser(user *models.User, db *pg.DB) error {
	return models.Save(user, db)
}

// UpdateUser the function update the User
func (u *DBUsers) UpdateUser(user *models.User, db *pg.DB) error {
	return models.Save(user, db)
}

// GetUser the function get user by msisdn
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
