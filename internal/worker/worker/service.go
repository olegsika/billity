package worker

import (
	"billity/common/models"
	"errors"
	"github.com/go-pg/pg"
)

type Service struct {
	callHistoryDB DBCallHistory
	userDB        UserDB
	db            *pg.DB
}

func New(callHistoryDB DBCallHistory, userDB UserDB, dbClient *pg.DB) *Service {
	return &Service{
		callHistoryDB: callHistoryDB,
		userDB:        userDB,
		db:            dbClient,
	}
}

type DBCallHistory interface {
	CreateCallHistory(history *models.CallHistory, db *pg.DB) error
}

type UserDB interface {
	GetUser(msisdn string, db *pg.DB) (models.User, error)
	UpdateBalance(user *models.User, db *pg.DB) error
}

func (s *Service) GetUser(sourceMsisdn string) (models.User, error) {
	return s.userDB.GetUser(sourceMsisdn, s.db)
}

func (s *Service) Create(history *models.CallHistory, balance float64) error {
	if history.TariffType == models.PrePaid {
		history.UserBalance = balance - history.RequestCost
	} else {
		history.UserBalance = balance
	}

	return s.callHistoryDB.CreateCallHistory(history, s.db)
}

func (s *Service) UpdateBalance(history *models.CallHistory, user models.User) error {
	if user.Id <= 0 {
		return errors.New("The user not found! ")
	}

	user.Balance -= history.RequestCost

	err := s.userDB.UpdateBalance(&user, s.db)

	if err != nil {
		return err
	}

	return nil
}
