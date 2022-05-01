package server

import "github.com/labstack/echo/v4"

// CustomBinder struct
type CustomBinder struct{}

// Bind tries to bind request into interface, and if it does then validate it
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	// You may use default binder
	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != nil && err != echo.ErrUnsupportedMediaType {
		return err
	}
	return c.Validate(i)
}
