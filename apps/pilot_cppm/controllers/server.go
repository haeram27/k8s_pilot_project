package controllers

import (
	"encoding/json"
	models "cppm/models"
	orm "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type ServerController struct {
	beego.Controller
}

type GetServerList struct {
	Header models.Header `json:"header"`
	Body models.ServerBody `json:"body"`
}

func (c *ServerController) Insert() {
	o := orm.NewOrm()

	svcIp := c.GetString("svcip")

	serverlist := new(models.ServerList)
	serverlist.Id = 1
	err := o.Read(serverlist)
	writeServer := new(models.ServerList)
	if err == orm.ErrNoRows {
		writeServer.ExternalIp = svcIp
		o.Insert(writeServer)
	} else {
		writeServer.Id = 1
		writeServer.ExternalIp = svcIp
		_ , err := o.Update(writeServer)
		if err != nil {
			panic(err)
		}
	}

	c.Redirect("../../", 200)
}

func (c *ServerController) Post() {
    o := orm.NewOrm()

    serverlist := new(models.ServerList)

	server := new(GetServerList)
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody), &server)
	if err != nil {
		panic(err)
	}
	serverlist.Id = 1
	writeServer := new(models.ServerList)
	err = o.Read(serverlist)
	if err == orm.ErrNoRows {
		writeServer.ExternalIp = server.Body.ExternalIp
		o.Insert(writeServer)
	} else {
		writeServer.Id = 1
		writeServer.ExternalIp = server.Body.ExternalIp
		_, err := o.Update(writeServer)
		if err != nil {
			panic(err)
		}
	}

    c.Redirect("/", 200)
}
