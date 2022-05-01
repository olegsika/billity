package report

import (
	"billity/common/models"
	"encoding/csv"
	"fmt"
	"github.com/go-pg/pg"
	"os"
	"strconv"
)

// Service the struct report service
type Service struct {
	reportDB DBReport
	dbClient *pg.DB
}

// New init report service
func New(reportDb DBReport, dbClient *pg.DB) *Service {
	return &Service{
		reportDB: reportDb,
		dbClient: dbClient,
	}
}

// DBReport interface for using call history table
type DBReport interface {
	GetDataForReport(msisdn string, from, to int, db *pg.DB) ([]models.CallHistory, error)
}

// GetData The function get data for report
func (s *Service) GetData(msisdn string, from, to int) ([]models.CallHistory, error) {
	callHistory, err := s.reportDB.GetDataForReport(msisdn, from, to, s.dbClient)

	if err != nil {
		return nil, err
	}

	return callHistory, nil
}

// GenerateCSV The function generate CSV file
func (s *Service) GenerateCSV(callHistory []models.CallHistory, msisdn string, from, to int) (*os.File, error) {
	fileName := generateFileName(msisdn, from, to)

	csvFile, err := os.Create(fileName)

	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(csvFile)

	defer func() {
		csvFile.Close()
	}()

	_ = writer.Write(getReportFields())

	for _, item := range callHistory {
		var row []string

		row = append(row, item.SourceMsisdn)
		row = append(row, item.DestinationMsisdn)
		row = append(row, string(item.Type))
		row = append(row, string(item.TariffType))
		row = append(row, fmt.Sprintf("%f", item.UserBalance))
		row = append(row, fmt.Sprintf("%f", item.RequestCost))
		row = append(row, fmt.Sprintf("%f", item.Tariff))
		row = append(row, strconv.Itoa(item.Duration))
		row = append(row, strconv.Itoa(item.CreatedAt))

		_ = writer.Write(row)
	}

	writer.Flush()

	return csvFile, nil
}

// getReportFields The function return report fields
func getReportFields() []string {
	return []string{
		"Source Msisdn",
		"Destination Msisdn",
		"Type",
		"Tariff Type",
		"User Balance",
		"Request Cost",
		"Tariff",
		"Duration",
		"CreatedAt",
	}
}

// generateFileName The function return generated file name
func generateFileName(msisdn string, from, to int) string {
	return fmt.Sprintf("call_history_%v_%v_%v.csv", msisdn, from, to)
}
