package bases

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"pilot_server/apps"
	"pilot_server/apps/funcs"
)

type Responses struct {
}

// REAT API 서버 응답 처리(Json 출력)
func (this *Responses) Response(c echo.Context, code int, data string) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	var output interface{}
	err := json.Unmarshal([]byte(data), &output)
	if nil != err {
		apps.Logs.Error("[Responses::Response] json parser error: ", funcs.OutputOmit(data))
		code, data = apps.GetHttpError(http.StatusInternalServerError, "json parser error")
		json.Unmarshal([]byte(data), &output)
	}

	c.Response().WriteHeader(code)

	enc := json.NewEncoder(c.Response())
	enc.SetIndent("", "  ")

	return enc.Encode(output)
}
