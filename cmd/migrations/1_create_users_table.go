package main

import "github.com/go-pg/migrations"

func init() {
	var tableName = "users"

	tableColumns := []map[string]string{
		{"id": "SERIAL PRIMARY KEY"},
		{"name": "varchar(255)"},
		{"balance": "float default 100"},
		{"msisdn": "varchar(14) default null unique"},
		{"tariff_type": "varchar(255) default null"},
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
