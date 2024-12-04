package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type envelope map[string]string

func (srv *Server) PostCount(e echo.Context) error {
	i2 := envelope{"count": "0"}
	err := e.Bind(&i2)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	err = srv.uc.FetchHelloMessage(i2["count"])
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "OK")
}
func (srv *Server) GetCount(e echo.Context) error {
	msg, err := srv.uc.GetHelloMessage()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}
