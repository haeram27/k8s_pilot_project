package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"pilot_agent/apps"
	"pilot_agent/controllers"
)

func Route(e *echo.Echo) {
	var homes = controllers.Homes{}
	var cmds = controllers.Cmds{}

	var dataRoutes = []struct {
		mode    string
		path    string
		handler echo.HandlerFunc
	}{
		{"GET", "/", homes.Get},
		{"GET", "/routes", homes.Routes},

		{"POST", "/cmds", cmds.Post},
	}

	for _, data := range dataRoutes {
		apps.Logs.Info(fmt.Sprintf("%s %s", data.mode, data.path))

		switch data.mode {
		case "GET":
			e.GET(data.path, data.handler)
		case "PUT":
			e.PUT(data.path, data.handler)
		case "POST":
			e.POST(data.path, data.handler)
		case "DELETE":
			e.DELETE(data.path, data.handler)
		default:
			apps.Logs.Error(fmt.Sprintf("unknown mode '%s'", data.mode))
		} // end of switch
	} // end of for
}
