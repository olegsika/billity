package service

import (
	"billity/cmd/api/request"
	"billity/internal/api/report"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Report struct {
	service *report.Service
}

func NewReport(reportService *report.Service, r *echo.Echo) {
	s := Report{
		service: reportService,
	}

	e := r.Group("/report")

	e.GET("/:msisdn", s.report)
}

func (r *Report) report(c echo.Context) error {
	reportReq, err := request.ReportRequest(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	errResp := request.ValidateReport(reportReq)

	if errResp != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errResp)
	}

	callHistory, err := r.service.GetData(reportReq.MSISDN, reportReq.From, reportReq.To)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	file, err := r.service.GenerateCSV(callHistory, reportReq.MSISDN, reportReq.From, reportReq.To)

	defer os.Remove(file.Name())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	return c.File(file.Name())
}
