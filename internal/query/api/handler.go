package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetQuery(e echo.Context) error {
	msg, err := srv.uc.FetchHelloMessage(e.QueryParam("name"))
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}
