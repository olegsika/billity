package postgres

import (
	"billity/common/models"
	"github.com/go-pg/pg"
)

// DBReport struct for ReportDB
type DBReport struct{}

// NewReportDB init ReportDB
func NewReportDB() *DBReport {
	return &DBReport{}
}

// GetDataForReport The function return data for report or error on fail
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
