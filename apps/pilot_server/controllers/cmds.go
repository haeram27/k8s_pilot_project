package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"pilot_server/apps"
	"pilot_server/apps/funcs"
	"pilot_server/controllers/bases"
	"pilot_server/controllers/cmds"
)

type Cmds struct {
	bases.Responses
}

func (this *Cmds) Post(c echo.Context) error {
	apps.Logs.Info("[Cmds::Post]")

	var jsonBody funcs.CmdJson
	if err := json.NewDecoder(c.Request().Body).Decode(&jsonBody); err != nil {
		code, data := apps.GetHttpError(http.StatusBadRequest, fmt.Sprintf("Invalid Request : %s", err.Error()))
		return this.Response(c, code, data)
	}
	apps.Logs.Info(jsonBody.Header.Type)

	// if err := cmds.Cmd_ProcList(jsonBody); err != nil {
	// 	code, data := apps.GetHttpError(http.StatusBadRequest, "Not Found Agent")
	// 	return this.Response(c, code, data)
	// }
	// code, data := apps.GetHttpError(http.StatusOK, "success")

	code, data := cmds.Cmd_ProcList(jsonBody)

	return this.Response(c, code, data)
}
