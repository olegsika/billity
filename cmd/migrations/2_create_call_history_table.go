package main

import "github.com/go-pg/migrations"

func init() {
	var tableName = "call_history"

	tableColumns := []map[string]string{
		{"id": "SERIAL PRIMARY KEY"},
		{"source_msisdn": "varchar(14) default null"},
		{"destination_msisdn": "varchar(14) default null"},
		{"type": "varchar(255) default null"},
		{"duration": "int default null"},
		{"tariff_type": "varchar(255) default null"},
		{"tariff": "float default 0"},
		{"request_cost": "float default null"},
		{"user_balance": "float default null"},
		{"created_at": "int NOT NULL DEFAULT date_part('epoch',CURRENT_TIMESTAMP)::int"},
		{"updated_at": "int NOT NULL DEFAULT date_part('epoch',CURRENT_TIMESTAMP)::int"},
		{"deleted_at": "int default null"},
	}

	_ = migrations.Register(func(db migrations.DB) error {

		return CreateTable(db, tableName, tableColumns)

	}, func(db migrations.DB) error {
		return DropTable(db, tableName)
	})
}
