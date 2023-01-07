package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"pilot_server/apps"
	"pilot_server/controllers/bases"
)

type Homes struct {
	bases.Responses
}

func (this *Homes) Get(c echo.Context) error {
	apps.Logs.Info("[Homes::Homes]")

	var data string = `{"apps":"pilot_server"}`

	return this.Response(c, http.StatusOK, data)
}
