package usage

import (
	"billity/common/models"
	"encoding/json"
	"errors"
	"github.com/go-pg/pg"
	"github.com/streadway/amqp"
	"sync"
)

type Service struct {
	usageDB         DBUsage
	db              *pg.DB
	rabbitmqChannel *amqp.Channel
	channelName     string
}

func New(usageDB DBUsage, dbClient *pg.DB, rabbitmqChannel *amqp.Channel, channelName string) *Service {
	return &Service{
		usageDB:         usageDB,
		db:              dbClient,
		rabbitmqChannel: rabbitmqChannel,
		channelName:     channelName,
	}
}

type DBUsage interface {
	GetBalance(msisdn string, db *pg.DB) (float64, error)
}

func (s *Service) ValidateBalance(history *models.CallHistory) (float64, error) {
	var wg sync.WaitGroup

	errChan := make(chan error)
	balanceChan := make(chan float64)
	wgDone := make(chan bool)
	requestCostChan := make(chan float64)

	wg.Add(2)

	go s.getUserBalance(history.SourceMsisdn, &wg, errChan, balanceChan)

	go s.calculateRequestCost(history.Type, history.Tariff, history.Duration, &wg, errChan, requestCostChan)

	go func(chan bool) {
		wg.Wait()
		wgDone <- true

	}(wgDone)

	var balance, requestCost float64

	for {
		select {
		case err := <-errChan:
			if err != nil {
				return 0, err
			}
		case resBalance := <-balanceChan:
			balance = resBalance
		case resRequestCost := <-requestCostChan:
			requestCost = resRequestCost
		case <-wgDone:
			close(balanceChan)
			close(requestCostChan)
			close(wgDone)
			close(errChan)
			if requestCost > (balance + 5) {

				return requestCost, errors.New("The request cost more than balance. ")
			}

			return requestCost, nil
		}
	}
}

func (s *Service) getUserBalance(msisdn string, wg *sync.WaitGroup, errChan chan error, balanceChan chan float64) {
	defer wg.Done()

	balance, err := s.usageDB.GetBalance(msisdn, s.db)

	if err != nil {
		errChan <- err
		return
	}

	balanceChan <- balance
}

func (s *Service) calculateRequestCost(callHistoryType models.CallHistoryType, tariff float64, callDuration int, wg *sync.WaitGroup, errChan chan error, requestCostChan chan float64) {
	defer wg.Done()

	var requestCost float64

	switch callHistoryType {
	case models.CallHistoryTypeCall:
		requestCost = tariff * float64(callDuration)
		requestCostChan <- requestCost
	case models.CallHistoryTypeSms:
		requestCost = tariff
		requestCostChan <- requestCost
	default:
		errChan <- errors.New("Not available call history type. ")
	}
}

func (s *Service) Publish(history *models.CallHistory) error {
	data, err := json.Marshal(history)

	if err != nil {
		return err
	}

	err = s.rabbitmqChannel.Publish(
		"",
		s.channelName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
