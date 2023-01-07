package controllers

import (
	"encoding/json"
	models "cppm/models"
	beego "github.com/beego/beego/v2/server/web"
	orm "github.com/beego/beego/v2/client/orm"
	"net/http"
	"bytes"
	"io/ioutil"
)

type NodeController struct {
	beego.Controller
}

type GetNodeList struct {
	Header models.Header `json:"header"`
	Body models.NodeBody `json:"body"`
}

func (c *NodeController) Get() {
	o := orm.NewOrm()

	var nodes []models.NodeList
	o.QueryTable("node_list").All(&nodes)
	c.Data["Nodes"] = nodes

    server := new(models.ServerList)
    server.Id = 1
    o.Read(server)
    c.Data["Server"] = server
	c.TplName = "nodes.tpl"
}

/*func (c *NodeController) Post() {
	o := orm.NewOrm()

	nodes := new(GetNodeList)
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody), &nodes)
	if err != nil {
		panic(err)
	}
	for _, node := range nodes.Body.NodeLists {
		nodelist := new(models.NodeList)
		nodelist.Name = node.Name
		nodelist.Ip = node.Ip
		o.Insert(nodelist)
	}

	c.Redirect("/nodes", 200)
}*/

type GatherProc struct {
	Cmds interface{} `json:"command_mode"`
}

type Commands struct {
	Mode string `json:"mode"`
}

type ResNodeList struct {
    Body []depth1 `json:"nodeList"`
}

func (c *NodeController) Gather() {
	o := orm.NewOrm()
	server := new(models.ServerList)
	server.Id = 1
	o.Read(server)

	reqSet := new(models.ReqSet)

	var reqHeader models.Header
	reqHeader.Version = "1"
	reqHeader.Type = "COMMAND_MODE"
	
	gatherProc := new(GatherProc)
	commands := new(Commands)
	commands.Mode = "GET_PODS_PROCLIST"
	gatherProc.Cmds = commands

	reqSet.Header = reqHeader
	reqSet.Body = gatherProc

	reqBody, err := json.MarshalIndent(reqSet, "", "    ")

	reqUri := "http://" + server.ExternalIp + "/cmds"

	resp, err := http.Post(reqUri, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		c.Redirect("/nodes", 302)
		panic(err)
	}

	pods := new(ResNodeList)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Redirect("/nodes", 302)
		panic(err)
	}

    err = json.Unmarshal([]byte(body), &pods)
    if err != nil {
        panic(err)
    }

	for _, node := range pods.Body {
    	nodeList := new(models.NodeList)
	    nodeList.NodeName = node.NodeName
	    err = o.Read(nodeList)
	    if err == orm.ErrNoRows {
	        nodeList.Timestamp = node.Timestamp
	        o.Insert(nodeList)
	    } else {
	        nodeList.Timestamp = node.Timestamp
	        o.Update(nodeList)
	        _, err := o.Raw("delete from proc_list where pod_name in (select pod_name from pod_list where node_name = ?)", nodeList.NodeName).Exec()
	        if err != nil {
	            panic(err)
	        }
	        _, err = o.Raw("delete from pod_list where node_name = ?;", nodeList.NodeName).Exec()
	        if err != nil {
	            panic(err)
	        }
	    }
		for _, pod := range node.PodList {
        	for _, proc := range pod.ProcList {
            	var newProc models.ProcList
	            newProc.Cmds = proc.Cmds
    	        newProc.Pid = proc.Pid
        	    newProc.Proc = proc.Proc
        	    newProc.PodName = pod.PodName
        	    o.Insert(&newProc)
        	}
        	var newPod models.PodList
        	newPod.ContainerId = pod.ContainerId
        	newPod.Namespace = pod.Namespace
        	newPod.PodName = pod.PodName
        	newPod.NodeName = nodeList.NodeName
       		o.Insert(&newPod)
    	}
	}
	
	c.Redirect("/nodes", 302)
}
