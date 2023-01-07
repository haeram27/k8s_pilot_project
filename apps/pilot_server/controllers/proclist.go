package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"pilot_server/apps"
	"pilot_server/apps/funcs"
	"pilot_server/controllers/bases"
	"pilot_server/controllers/db"
)

type ProcList struct {
	bases.Responses
}

func (this *ProcList) GetPodProcList(c echo.Context) error {
	apps.Logs.Info("[ProcList::GetPodProcList]")

	ctx := context.Background()
	client := db.GetClient()
	defer client.Close()

	nodes, err := client.Node.Query().All(ctx)

	if err != nil {
		apps.Logs.Error("failed creating node: ", err)
		code, data := apps.GetHttpError(http.StatusInternalServerError, fmt.Sprintf("failed reading node : %s", err.Error()))
		return this.Response(c, code, data)
	}

	var nodeList funcs.NodeList
	nodeList.Mode = "GET_PODS_PROCLIST"

	for _, n := range nodes {
		apps.Logs.Info("node >>>> ", n)

		var jsonData funcs.PodProcList

		jsonData.NodeName = n.ID
		jsonData.TimeStamp = n.Timestamp

		if err := json.Unmarshal([]byte(n.PodInfo), &jsonData.PodList); err != nil {
			apps.Logs.Error("failed unmarshal")
			continue
		}

		// apps.Logs.Info(jsonData)

		nodeList.NodeList = append(nodeList.NodeList, jsonData)
	}

	return this.Response(c, http.StatusOK, funcs.JsonDumps(nodeList))
}
