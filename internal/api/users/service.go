package users

import (
	"billity/common/models"
	"errors"
	"github.com/go-pg/pg"
	"time"
)

// Service the struct usage service
type Service struct {
	userDB DBUsers
	db     *pg.DB
}

// New init users service
func New(userDB DBUsers, dbClient *pg.DB) *Service {
	return &Service{
		userDB: userDB,
		db:     dbClient,
	}
}

type DBUsers interface {
	CreateUser(user *models.User, db *pg.DB) error
	UpdateUser(user *models.User, db *pg.DB) error
	GetUser(msisdn string, isForUpdate bool, db *pg.DB) (models.User, error)
}

// Create the function create the User
func (s *Service) Create(user *models.User) error {
	return s.userDB.CreateUser(user, s.db)
}

// Update the function update the User
func (s *Service) Update(user *models.User) error {

	originUser, err := s.userDB.GetUser(user.MSISDN, true, s.db)

	if err != nil {
		return err
	}

	if originUser.Id <= 0 {
		return errors.New("The user does not exists or deleted! ")
	}

	user.Id = originUser.Id
	user.CreatedAt = originUser.CreatedAt
	user.UpdatedAt = originUser.UpdatedAt
	user.DeletedAt = originUser.DeletedAt

	return s.userDB.UpdateUser(user, s.db)
}

// Delete the function delete the User
func (s *Service) Delete(msisdn string) error {
	user, err := s.userDB.GetUser(msisdn, false, s.db)

	if err != nil {
		return err
	}

	if user.Id <= 0 {
		return errors.New("The user does not exists or deleted ")
	}

	user.DeletedAt = int(time.Now().UTC().Unix())

	return s.userDB.UpdateUser(&user, s.db)
}

// AddBalance the function add balance for the User
func (s *Service) AddBalance(balance float64, msisdn string) error {
	user, err := s.userDB.GetUser(msisdn, false, s.db)

	if err != nil {
		return err
	}

	if user.Id <= 0 {
		return errors.New("Id can not be empty! ")
	}

	user.Balance += balance

	return s.userDB.UpdateUser(&user, s.db)
}
