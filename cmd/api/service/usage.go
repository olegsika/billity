package service

import (
	"billity/cmd/api/request"
	"billity/common/models"
	"billity/internal/api/usage"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Usage the service for manage user balance
type Usage struct {
	service *usage.Service
}

// NewUsage init the service
func NewUsage(usageService *usage.Service, r *echo.Echo) {
	s := Usage{
		service: usageService,
	}

	e := r.Group("/usage")

	e.POST("", s.usage)
}

// usage the function get request, validate and sent request to worker
func (u *Usage) usage(c echo.Context) error {
	callHistory, err := request.UsageRequest(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	errResp := request.ValidateUsage(callHistory)

	if errResp != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	callHistory.RequestCost, err = u.service.ValidateBalance(callHistory)

	if err != nil && callHistory.TariffType == models.PrePaid {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	err = u.service.Publish(callHistory)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "Balance Updated.")
}
