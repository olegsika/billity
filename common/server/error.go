package server

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type customErrHandler struct {
	e *echo.Echo
}

var validationErrors = map[string]string{
	"required": "required, but was not received",
	"min":      "value or length is less than allowed",
	"max":      "value or length is bigger than allowed",
}

func getVldErrorMsg(s string) string {
	if v, ok := validationErrors[s]; ok {
		return v
	}
	return " failed on " + s + " validation"
}

func (ce *customErrHandler) handler(err error, c echo.Context) {
	var (
		code = http.StatusUnprocessableEntity
		msg  interface{}
	)

	type resp struct {
		Status  bool        `json:"status"`
		Message interface{} `json:"error"`
	}

	switch er := err.(type) {
	case *echo.HTTPError:
		code = er.Code
		msg = er.Message
		if er.Internal != nil {
			msg = fmt.Sprintf("%v, %v", err, er.Internal)
		}
	case validator.ValidationErrors:
		errMsg := make(map[string]interface{})
		for _, v := range er {
			errMsg[v.Field()] = getVldErrorMsg(v.ActualTag())
		}
		msg = resp{Message: errMsg}
		code = http.StatusBadRequest
	case pg.Error:
		code = http.StatusNoContent
	default:
		msg = err.Error()
	}
	if _, ok := msg.(string); ok {
		msg = resp{Message: msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			if err = c.NoContent(code); err != nil {
				ce.e.Logger.Error(err)
				return
			}
		}
		if code == 204 {
			if err = c.NoContent(code); err != nil {
				ce.e.Logger.Error(err)
			}
			return
		}
		if err = c.JSON(code, msg); err != nil {
			ce.e.Logger.Error(err)
		}
	}

}
