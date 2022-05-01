package main

import (
	"billity/cmd/api/service"
	"billity/common/config"
	"billity/common/db"
	"billity/common/server"
	"billity/common/utils"
	"billity/internal/api/report"
	reportDB "billity/internal/api/report/platform/postgres"
	"billity/internal/api/usage"
	usageDB "billity/internal/api/usage/platform/postgres"
	"billity/internal/api/users"
	userDB "billity/internal/api/users/platform/postgres"
	"flag"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
)

// main The function run API microservice
func main() {
	cfgPath := flag.String("p", "./cmd/api/config/config.yaml", "Path to config file")
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

	q, err := ch.QueueDeclare(
		cfg.RabbitMQ.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.CheckErr(err)

	fmt.Println(q)

	e := server.New()

	addServices(e, dbClient, ch, cfg.RabbitMQ.QueueName)

	server.Start(e, cfg.Server)
}

// addServices the function init DB and run all the entrypoint for ms
func addServices(e *echo.Echo, dbClient *pg.DB, rabbitmqChannel *amqp.Channel, queueName string) {

	// Init Users Service
	dbUsers := userDB.NewUsersDB()
	usersService := users.New(dbUsers, dbClient)
	service.NewUser(usersService, e)

	// Init Usage Service
	dbUsage := usageDB.NewUsageDB()
	usageService := usage.New(dbUsage, dbClient, rabbitmqChannel, queueName)
	service.NewUsage(usageService, e)

	// Init Report Service
	dbReport := reportDB.NewReportDB()
	reportService := report.New(dbReport, dbClient)
	service.NewReport(reportService, e)
}
