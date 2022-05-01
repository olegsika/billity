package request

import (
	"billity/common/models"
	"billity/common/utils"
	"github.com/labstack/echo/v4"
	"strings"
)

func UsageRequest(c echo.Context) (*models.CallHistory, error) {
	callHistory := new(models.CallHistory)

	if err := c.Bind(callHistory); err != nil {
		return nil, err
	}

	return callHistory, nil
}

func ValidateUsage(callHistory *models.CallHistory) map[string]string {
	errResp := make(map[string]string)

	if !utils.InListString(string(callHistory.TariffType), []string{string(models.PrePaid), string(models.PostPaid)}) {
		errResp["tariff_type"] = "Tariff Type is invalid."
	}
	if !utils.InListString(string(callHistory.Type), []string{string(models.CallHistoryTypeCall), string(models.CallHistoryTypeSms)}) {
		errResp["call_type"] = "Call Type is invalid."
	}

	if callHistory.SourceMsisdn == "" {
		errResp["source_msisdn"] = "MSISDN is invalid."
	}

	if !strings.HasPrefix(callHistory.SourceMsisdn, "+") {
		errResp["source_msisdn_prefix"] = "MSISDN will begin with \"+\""
	}

	if len(callHistory.SourceMsisdn) > 14 {
		errResp["source_msisdn_length"] = "The msisdn too long."
	}

	if callHistory.DestinationMsisdn == "" {
		errResp["destination_msisdn"] = "MSISDN is invalid."
	}

	if !strings.HasPrefix(callHistory.DestinationMsisdn, "+") {
		errResp["destination_msisdn_prefix"] = "MSISDN will begin with \"+\""
	}

	if len(callHistory.DestinationMsisdn) > 14 {
		errResp["destination_msisdn_length"] = "The msisdn too long."
	}

	if len(errResp) > 0 {
		return errResp
	}

	return nil
}
