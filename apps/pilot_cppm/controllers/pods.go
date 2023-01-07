package controllers

import (
	"encoding/json"
	models "cppm/models"
	beego "github.com/beego/beego/v2/server/web"
	orm "github.com/beego/beego/v2/client/orm"

	"fmt"
)

type PodController struct {
	beego.Controller
}

type GetPodList struct {
    Header models.Header `json:"header"`
    Body depth1 `json:"body"`
}

type depth1 struct {
	NodeName string `json:"nodeName"`
	Timestamp string `json:"timestamp"`
	PodList []depth2 `json:"podList"`
}

type depth2 struct {
	ContainerId string `json:"containerID"`
	Namespace string `json:"namespace"`
	PodName string `json:"podName"`
	ProcList []depth3 `json:"procList"`
}

type depth3 struct {
	Cmds string `json:"cmds"`
	Pid uint64 `json:"pid"`
	Proc string `json:"proc"`
}

func (c *PodController) Post() {
	o := orm.NewOrm()

	pods := new(GetPodList)
	pods.Body = depth1{}
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody), &pods)
	if err != nil {
		panic(err)
	}

	nodeList := new(models.NodeList)
	nodeList.NodeName = pods.Body.NodeName
	err = o.Read(nodeList)
	if err == orm.ErrNoRows {
		nodeList.Timestamp = pods.Body.Timestamp
		o.Insert(nodeList)
	} else {
		nodeList.Timestamp = pods.Body.Timestamp
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

	for _, pod := range pods.Body.PodList {
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

	c.Redirect("/nodes", 201)
}

type JoinReceiver struct {
	NodeName string
	ContainerID string
	Namespace string
	PodName string
	Cmds string
	Pid uint64
	Proc string
}

func (c *PodController) Get() {
	o := orm.NewOrm()

	server := new(models.ServerList)
	server.Id = 1
	o.Read(server)
	c.Data["Server"] = server

	nodeName := c.Ctx.Input.Param(":name")
	//query := fmt.Sprintf("select * from proc_list where (select pod_name from pod_list where node_name='%s')=pod_name;", nodeName)
	//query := fmt.Sprintf("select b.container_id, b.namespace, a.cmds, a.pid, a.proc, a.pod_name from proc_list as a inner join pod_list as b where a.pod_name = b.pod_name and b.node_name=%s", nodeName)
	query := fmt.Sprintf("select b.node_name, b.container_id, b.namespace, a.cmds, a.pid, a.proc, a.pod_name from proc_list as a inner join pod_list as b where a.pod_name=b.pod_name and b.node_name='%s';", nodeName)

	//var procs []models.ProcList
	var procs []JoinReceiver

	_, err := o.Raw(query).QueryRows(&procs)
	if err != nil{
		panic(err)
	}

	c.Data["Procs"] = procs
	if len(procs) > 0 {
		c.Data["NodeName"] = nodeName
	}
	c.TplName = "pods.tpl"
}

/*func (c *PodController) RegisterPod() {
	o := orm.NewOrm()

	pods := new(GetPodList)
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody), &pods)
	fmt.Println(pods)
	if err != nil {
		panic(err)
	}
	
	for _, pod := range pods.Body.PodLists {
		podlist := new(models.PodList)
		podlist.NodeName = pod.NodeName
		podlist.PodName = pod.PodName
		o.Insert(podlist)
		fmt.Println(podlist)
	}

	c.Redirect("/nodes", 200)
}*/
