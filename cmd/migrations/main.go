package main

import (
	"billity/common/config"
	hdb "billity/common/db"
	"billity/common/utils"
	"flag"
	"fmt"
	"github.com/go-pg/migrations"
	"os"
)

func main() {
	cfgPath := flag.String("p", "./cmd/migrations/config/config.yaml", "Path to config file")
	flag.Parse()

	// Load Configs from file
	cfg, err := config.LoadConfigs(*cfgPath)
	utils.CheckErr(err)

	db, err := hdb.NewGoPG(cfg.DbPSN)
	utils.CheckErr(err)
	defer db.Close()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)

	if err != nil {
		exitf(err.Error())
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func CreateTable(db migrations.DB, tableName string, tableColumns []map[string]string) error {
	var separate string = ""

	var sql = `CREATE TABLE IF NOT EXISTS ` + tableName + ` (`

	for _, newColumnInfo := range tableColumns {

		for columnName, columnDefinition := range newColumnInfo {
			sql += separate + columnName + " " + columnDefinition

			if separate == "" {
				separate = ", "
			}
		}
	}

	sql += `);`

	_, err := db.Exec(sql)

	return err
}

func DropTable(db migrations.DB, tableName string) error {
	var sql = `DROP TABLE ` + tableName + `;`

	_, err := db.Exec(sql)

	return err
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
