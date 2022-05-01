package server

import (
	"billity/common/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// New init new server
func New(mws ...echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()

	for _, v := range mws {
		e.Use(v)
	}
	e.GET("/health/", healthCheck)
	e.Validator = &CustomValidator{V: validator.New()}
	custErr := &customErrHandler{e: e}
	e.HTTPErrorHandler = custErr.handler
	e.Binder = &CustomBinder{}
	return e
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "UP")
}

// Start starts echo server
func Start(e *echo.Echo, cfg config.Server) {
	e.Server.Addr = cfg.Port
	e.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Second
	e.Debug = cfg.Debug
	e.Logger.Fatal(e.Start(cfg.Port))
}
