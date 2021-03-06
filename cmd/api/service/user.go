package service

import (
	"billity/cmd/api/request"
	"billity/internal/api/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

// User the service for manage user
type User struct {
	service *users.Service
}

// NewUser init the service
func NewUser(userService *users.Service, r *echo.Echo) {
	s := User{
		service: userService,
	}

	e := r.Group("/users")

	e.POST("", s.create)

	e.PUT("/add-balance/:msisdn", s.addBalance)
	e.PUT("/:msisdn", s.update)
	e.DELETE("/:msisdn", s.delete)
}

// create The function will create the user
func (u *User) create(c echo.Context) error {
	user, err := request.UserRequest(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	errResp := request.Validate(user)

	if errResp != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errResp)
	}

	err = u.service.Create(user)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// update The function will update the user
func (u *User) update(c echo.Context) error {
	user, err := request.UserRequest(c)

	user.MSISDN = c.Param("msisdn")

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	errResp := request.Validate(user)

	if errResp != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errResp)
	}

	err = u.service.Update(user)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// delete The function will delete the user
func (u *User) delete(c echo.Context) error {
	msisdn := c.Param("msisdn")

	if msisdn == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "The msisdn can not be empty")
	}

	err := u.service.Delete(msisdn)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}

// addBalance The function will add balance for the user
func (u *User) addBalance(c echo.Context) error {
	msisdn := c.Param("msisdn")

	if msisdn == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "The msisdn can not be empty")
	}

	user, err := request.UserRequest(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if user.Balance <= 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "The balance is empty or negative.")
	}

	err = u.service.AddBalance(user.Balance, msisdn)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "Balance Updated.")
}
