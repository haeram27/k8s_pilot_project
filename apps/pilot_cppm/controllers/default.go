package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	orm "github.com/beego/beego/v2/client/orm"
	models "cppm/models"
	"net/http"
	"encoding/json"
	"bytes"
	"strconv"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*o := orm.NewOrm()

	server := new(models.ServerList)
	server.Id = 1
	o.Read(server)

	c.Data["Server"] = server*/

	c.Redirect("/nodes", 302)
}

type RegisterServer struct {
	RegisterCPPM interface{} `json:"REGISTER_CPPM"`
}

type Ip struct {
	CPPMIp string `json:"ip"`
}

func (c *MainController) Insert() {
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

	reqUri := "http://" + svcIp + "/cmds"
	
	reqSet := new(models.ReqSet)

	var reqHeader models.Header
	reqHeader.Version = "1"
	reqHeader.Type = "REGISTER_CPPM"

	registerServer := new(RegisterServer)
	ip := new(Ip)
	ip.CPPMIp = c.Ctx.Input.Host() + ":" + strconv.Itoa(c.Ctx.Input.Port())
	registerServer.RegisterCPPM = ip

	reqSet.Header = reqHeader
	reqSet.Body = registerServer

	reqBody, err := json.MarshalIndent(reqSet, "", "    ")

	resp, err := http.Post(reqUri, "application/json", bytes.NewBuffer(reqBody))
	_ = resp

	c.Redirect("/", 302)
}

func (c *MainController) Cmds() {
	fmt.Println(string(c.Ctx.Input.RequestBody))
	c.Redirect("/", 404)
}
