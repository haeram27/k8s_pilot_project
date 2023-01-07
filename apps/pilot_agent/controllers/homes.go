package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"net/http"

	"pilot_agent/apps"
	"pilot_agent/controllers/bases"
)

type Homes struct {
	bases.Responses
}

func (this *Homes) Get(c echo.Context) error {
	apps.Logs.Info("[Homes::Homes]")

	data := fmt.Sprintf(
		`{"apps":"pilot_agent","version":"%s"}`, apps.Conf.Agent.Version)

	return this.Response(c, http.StatusOK, data)
}

func (this *Homes) Routes(c echo.Context) error {
	apps.Logs.Info("[Homes::Routes]")

	echo := apps.GetEcho()
	json_data, _ := json.Marshal(echo.Routes())
	data := fmt.Sprintf(`{"routes":%s}`, string(json_data))

	return this.Response(c, http.StatusOK, data)
}
