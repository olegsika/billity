package request

import (
	"billity/common/models"
	"billity/common/utils"
	"github.com/labstack/echo/v4"
	"strings"
)

func UserRequest(c echo.Context) (*models.User, error) {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return nil, err
	}

	return user, nil
}

func Validate(user *models.User) map[string]string {
	errResp := make(map[string]string)

	if user.Name == "" {
		errResp["name"] = "User Name is empty"
	}

	if user.Balance < 0 {
		errResp["balance"] = "You can not create User with a negative balance."
	}

	if !utils.InListString(string(user.TariffType), []string{string(models.PrePaid), string(models.PostPaid)}) {
		errResp["tariff_type"] = "Tariff Type is invalid."
	}

	if user.MSISDN == "" {
		errResp["msisdn"] = "MSISDN is invalid."
	}

	if !strings.HasPrefix(user.MSISDN, "+") {
		errResp["msisdn_prefix"] = "MSISDN will begin with \"+\""
	}

	if len(user.MSISDN) > 14 {
		errResp["msisdn_length"] = "The msisdn too long."
	}

	if len(errResp) > 0 {
		return errResp
	}

	return nil
}
