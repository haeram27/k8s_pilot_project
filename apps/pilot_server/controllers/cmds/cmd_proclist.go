package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"pilot_server/apps"
	"pilot_server/apps/funcs"

	"pilot_server/controllers/db"

	_ "github.com/lib/pq"
)

func parseBody(respBody []byte) error {
	apps.Logs.Info("[Cmds::parseBody]")

	var jsonData funcs.PodProcList
	if err := json.Unmarshal(respBody, &jsonData); err != nil {
		return err
	}

	if err := db.CreateNode(jsonData); err != nil {
		return err
	}

	return nil
}

func postRequest(url string, requestBody funcs.CmdJson, podlist *funcs.PodProcList) (int, string) {
	apps.Logs.Info("[Cmds::postRequest]")

	body, err := json.Marshal(requestBody)
	if err != nil {
		return apps.GetHttpError(http.StatusBadRequest, "Failed to marshal request body")
	}

	buff := bytes.NewBuffer(body)
	apps.Logs.Info(buff)

	resp, err := http.Post(url, "application/json", buff)
	if err != nil {
		return apps.GetHttpError(http.StatusBadRequest, "Failed to post reuqest to agent")
	}
	defer resp.Body.Close()

	// response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return apps.GetHttpError(http.StatusBadRequest, "Failed to read response")
	}

	apps.Logs.Info(string(respBody))

	if err := json.Unmarshal(respBody, podlist); err != nil {
		return apps.GetHttpError(http.StatusBadRequest, "Failed to unmarshal")
	}

	if err := db.CreateNode(*podlist); err != nil {
		return apps.GetHttpError(http.StatusBadRequest, "Failed to create node db")
	}

	return http.StatusOK, "success"
}

func Cmd_ProcList(jsonBody funcs.CmdJson) (int, string) {
	apps.Logs.Info("[Cmds::Cmd_ProcList]")

	apps.Logs.Info(apps.Conf.Agent.Addr)
	apps.Logs.Info(apps.Conf.Server.Port)

	_, addrs, err := net.LookupSRV("", "tcp", apps.Conf.Agent.Addr)
	if err != nil {
		return apps.GetHttpError(http.StatusBadRequest, "failed to get agent address")
	}

	var nodeList funcs.NodeList
	nodeList.Mode = "GET_PODS_PROCLIST_1"

	for _, addr := range addrs {
		agentAddr := addr.Target[:len(addr.Target)-1]

		apps.Logs.Info(fmt.Sprintf("[Daemonset ADDR] %s:%d", agentAddr, addr.Port))

		var jsonData funcs.PodProcList

		url := fmt.Sprintf("http://%s:%d/cmds", agentAddr, addr.Port)
		code, _ := postRequest(url, jsonBody, &jsonData)
		if code != http.StatusOK {
			continue
		}
		nodeList.NodeList = append(nodeList.NodeList, jsonData)

		apps.Logs.Info(nodeList)

	}

	return http.StatusOK, funcs.JsonDumps(nodeList)
}
