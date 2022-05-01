package main

import (
	"billity/cmd/worker/service"
	"billity/common/config"
	"billity/common/db"
	"billity/common/utils"
	"billity/internal/worker/worker"
	"billity/internal/worker/worker/platform/postgres"
	"flag"
	"github.com/go-pg/pg"
	"github.com/streadway/amqp"
)

// main The function run Worker microservice
func main() {
	cfgPath := flag.String("p", "./cmd/worker/config/config.yaml", "Path to config file")
	flag.Parse()

	// Load Configs from file
	cfg, err := config.LoadConfigs(*cfgPath)
	utils.CheckErr(err)

	// Init Database
	dbClient, err := db.NewGoPG(cfg.DbPSN)
	utils.CheckErr(err)
	defer dbClient.Close()

	rabbitmqConn, err := amqp.Dial(cfg.RabbitMQ.RabbitMQConnUrl)
	utils.CheckErr(err)
	defer rabbitmqConn.Close()

	ch, err := rabbitmqConn.Channel()
	defer ch.Close()
	utils.CheckErr(err)

	addServices(dbClient, ch, cfg.RabbitMQ.QueueName)
}

// addServices the function init DB and run all the entrypoint for ms
func addServices(dbClient *pg.DB, rabbitmqChannel *amqp.Channel, queueName string) {
	dbCallHistory := postgres.NewCallHistoryDB()
	dbUsers := postgres.NewUserDB()
	workerService := worker.New(dbCallHistory, dbUsers, dbClient)
	service.NewWorker(workerService, rabbitmqChannel, queueName)
}
