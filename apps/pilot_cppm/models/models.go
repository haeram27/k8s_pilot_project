package models

import (
	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"
	orm "github.com/beego/beego/v2/client/orm"
)

type ReqSet struct {
	Header Header `json:"header"`
	Body interface{} `json:"body"`
}

type Header struct {
	Version string `json:"version"`
	Type string `json:"type"`
}

type ServerBody struct {
	ExternalIp string `json:"external_ip"`
}

type ServerList struct {
    Id int `orm:"auto"`
    ExternalIp string `orm:"size(30)"`
}

type NodeBody struct {
	NodeLists []NodeList `json:"nodes"`
}

type NodeList struct {
    //Id int `orm:"auto"`
	NodeName string `json:"nodeName" orm:"pk"`
	Timestamp string `json:"timestamp" orm:"size(30)"`
	//Pods []*PodList `orm:"reverse(many)"`
}

type PodBody struct {
	PodLists []PodList `json:"pods"`
}

type PodList struct {
	//Id int `orm:"auto"`
	//Node *NodeList `orm:"rel(fk)"`
	ContainerId string `json:"containerID" orm:"size(50)"`
	Namespace string `json:"namespace" orm:"size(30)"`
	PodName string `json:"podName" orm:"pk"`
	NodeName string `orm:"size(30)"`
	//Procs []*ProcList `orm:"reverse(many)"`
}

type ProcList struct {
	Id int `orm:"auto"`
	//Pod *PodList `orm:"rel(fk)"`
	Cmds string `json:"cmds" orm:"size(50)"`
	Pid uint64 `json:"pid" orm:"size(10)"`
	Proc string `json:"proc" orm:"size(50)"`
	PodName string `orm:"size(30)"`
}

func init() {
    orm.RegisterDriver("sqlite3", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "data.db")
    orm.RegisterModel(new(ServerList), new(NodeList), new(PodList), new(ProcList))

	orm.RunSyncdb("default", false, true)
}
