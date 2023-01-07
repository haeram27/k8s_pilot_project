package controllers

import (
	"github.com/labstack/echo/v4"

	"pilot_agent/apps"
	"pilot_agent/controllers/bases"
	"pilot_agent/controllers/cmds"
)

type Cmds struct {
	bases.Responses

	cmds.CmdsImpl
}

func (this *Cmds) Post(c echo.Context) error {
	apps.Logs.Info("[Cmds::Post]")

	code, data := this.OnCommand(c)

	return this.Response(c, code, data)
}
