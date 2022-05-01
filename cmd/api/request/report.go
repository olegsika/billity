package request

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type ReportReq struct {
	From   int    `json:"from"`
	To     int    `json:"to"`
	MSISDN string `json:"msisdn"`
}

// ReportRequest the function Bind request to model
func ReportRequest(c echo.Context) (*ReportReq, error) {
	req := new(ReportReq)

	var err error

	req.MSISDN = c.Param("msisdn")
	req.From, err = strconv.Atoi(c.QueryParam("from"))
	if err != nil {
		return nil, err
	}

	req.To, err = strconv.Atoi(c.QueryParam("to"))
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ValidateReport the function validate request
func ValidateReport(req *ReportReq) map[string]string {
	errResp := make(map[string]string)

	if req.From > req.To {
		errResp["from"] = "The FROM can not be greater than TO"
	}

	if req.MSISDN == "" {
		errResp["msisdn"] = "The msisdn can not be empty"
	}

	if !strings.HasPrefix(req.MSISDN, "+") {
		errResp["msisdn_prefix"] = "MSISDN will begin with \"+\""
	}

	if len(req.MSISDN) > 14 {
		errResp["msisdn_length"] = "The msisdn too long."
	}

	if len(errResp) > 0 {
		return errResp
	}

	return nil
}
