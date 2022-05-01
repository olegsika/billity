package service

import (
	"billity/cmd/worker/request"
	"billity/common/models"
	"billity/internal/worker/worker"
	"fmt"
	"github.com/streadway/amqp"
)

type Worker struct {
	service   *worker.Service
	channel   *amqp.Channel
	queueName string
}

func NewWorker(service *worker.Service, channel *amqp.Channel, queueName string) {
	s := Worker{
		service:   service,
		channel:   channel,
		queueName: queueName,
	}

	s.Run()
}

func (w *Worker) Run() {
	msqs, err := w.channel.Consume(
		w.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msqs {
			fmt.Printf("Received new message %s\n", d.Body)
			err := w.SaveHistory(d.Body)

			if err != nil {
				continue
			}

		}
	}()

	fmt.Println("Successfully connected to RabbitMQ instance")
	fmt.Println("Waiting for messages")

	<-forever
}

func (w *Worker) SaveHistory(data []byte) error {
	callHistory, err := request.CallHistoryRequest(data)

	if err != nil {
		return err
	}

	user, err := w.service.GetUser(callHistory.SourceMsisdn)

	if err != nil {
		return err
	}

	err = w.service.Create(callHistory, user.Balance)

	if err != nil {
		return err
	}

	if callHistory.TariffType == models.PrePaid {
		err = w.service.UpdateBalance(callHistory, user)

		if err != nil {
			return err
		}
	}

	return nil
}
