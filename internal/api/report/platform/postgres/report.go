package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

type DBReport struct{}

func NewReportDB() *DBReport {
	return &DBReport{}
}

func (r *DBReport) GetDataForReport(msisdn string, from, to int, db *pg.DB) ([]models.CallHistory, error) {
	callHistories := []models.CallHistory{}

	err := db.Model(&callHistories).
		Where("source_msisdn = ?", msisdn).
		Where("created_at between ? and ?", from, to).
		Where("deleted_at is null").
		Order("created_at desc").
		Select()

	if err != nil {
		return nil, err
	}

	return callHistories, nil
}
