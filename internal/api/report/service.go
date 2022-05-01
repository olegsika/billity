package report

import (
	"billity/common/models"
	"encoding/csv"
	"fmt"
	"github.com/go-pg/pg"
	"os"
	"strconv"
)

type Service struct {
	reportDB DBReport
	dbClient *pg.DB
}

func New(reportDb DBReport, dbClient *pg.DB) *Service {
	return &Service{
		reportDB: reportDb,
		dbClient: dbClient,
	}
}

type DBReport interface {
	GetDataForReport(msisdn string, from, to int, db *pg.DB) ([]models.CallHistory, error)
}

type DBUser interface {
	GetUserBalance(msisdn, db *pg.DB) (float64, error)
}

func (s *Service) GetData(msisdn string, from, to int) ([]models.CallHistory, error) {
	callHistory, err := s.reportDB.GetDataForReport(msisdn, from, to, s.dbClient)

	if err != nil {
		return nil, err
	}

	return callHistory, nil
}

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
func generateFileName(msisdn string, from, to int) string {
	return fmt.Sprintf("call_history_%v_%v_%v.csv", msisdn, from, to)
}
