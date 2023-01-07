package cmds

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"pilot_agent/apps"
	"pilot_agent/apps/funcs"
)

type CmdsImpl struct {
}

func (this *CmdsImpl) OnCommand(c echo.Context) (int, string) {
	var json_data map[string]interface{}

	err := json.NewDecoder(c.Request().Body).Decode(&json_data)
	if err != nil {
		apps.Logs.Error(fmt.Sprintf("[CmdsImpl::OnCommand] json.NewDecoder failed(%v)", err))
		return apps.GetHttpError(http.StatusBadRequest, "Invalid JSON data")
	}

	/*
		{
			"body":{
				"command_mode":{
					"mode":"GET_PODS_PROCLIST"
				}
			}
		}
	*/
	value, stype := funcs.JsonValue(json_data, "body", "command_mode", "mode")
	if "string" != stype {
		apps.Logs.Error(fmt.Sprintf("[CmdsImpl::OnCommand] funcs.JsonValue failed"))
		return apps.GetHttpError(http.StatusBadRequest, "Field not found in the JSON data")
	}

	mode := value.(string)
	code, data := apps.GetHttpError(http.StatusBadRequest, "Invalid mode")

	switch mode {
	case "GET_PODS_PROCLIST":
		code, data = this.cmdGetPodsProclist()
	}

	apps.Logs.Info(fmt.Sprintf("[CmdsImpl::OnCommand] mode=%s,code=%d", mode, code))

	return code, data
}

func (this *CmdsImpl) cmdGetPodsProclist() (int, string) {
	return getPodsProclist()
}
