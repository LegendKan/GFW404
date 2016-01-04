package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sync"
	"strconv"
	"ssm/models"
)

var lock sync.Mutex

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	//c.Data["Address"]
	//"ssm/controllers/unixsocket"
	//c.Ctx.WriteString(MasterAddr + ":" + MasterPort + ":" + passAuth + "\n" + unixsocket.SocketTest())
	var t models.Account
	t.Id=1
	t.Port=9009
	t.Containerid="123456"
	t.Password="111111"
	a:=make([]models.Account,1)
	a=append(a,t)
	SyncContainers(a)
	c.Ctx.WriteString("Success")
}

type AddController struct {
	beego.Controller
}

func (c *AddController) Get() {
	port := c.GetString("port")
	pass := c.GetString("pass")
	if len(pass) <= 0 {
		pass = randSeq(8)
	}
	if len(port)<=0{
		lock.Lock()
		port=strconv.Itoa(MaxPort)
		MaxPort++
		lock.Unlock()
	}
	r := make(map[string]interface{})
	ss := make(map[string]string)
	if ret, ok := AddContainer(port, pass); ok {
		r["status"] = true
		ss["server"] = serverAddr
		ss["password"] = pass
		ss["containerid"] = ret
		if len(port) <= 0 {
			//
			ssconfig := InspectContainer(ret)
			ssjson := make(map[string]interface{})
			json.Unmarshal([]byte(ssconfig), &ssjson)
			//tmp := ssjson.(map[string]interface{})
			tmp1 := ssjson["NetworkSettings"]
			tmp2 := tmp1.(map[string]interface{})
			tmp3 := tmp2["Ports"]
			tmp4 := tmp3.(map[string]interface{})
			tmp5 := tmp4[innerPort+"/tcp"].([]interface{})
			tmp6 := tmp5[0].(map[string]interface{})
			//ss["port"] = tmp["HostConfig"]["PortBindings"][innerPort+"/tcp"][0]["HostPort"]
			ss["port"] = tmp6["HostPort"].(string)

		} else {
			ss["port"] = port
		}
		r["data"] = ss
	} else {
		r["status"] = false
		r["data"] = "Error"
	}
	b, _ := json.Marshal(r)
	c.Ctx.WriteString(string(b))
}

type DeleteController struct {
	beego.Controller
}

func (c *DeleteController) Get() {
	cid := c.GetString("cid")
	ret := make(map[string]interface{})
	if len(cid) <= 0 {
		ret["status"] = false
		ret["data"] = "No Container ID"
	} else {
		if DeleteContainer(cid) {
			ret["status"] = true
		} else {
			ret["status"] = false
			ret["data"] = "Error happened while deleting"
		}
	}
	b, _ := json.Marshal(ret)
	c.Ctx.WriteString(string(b))
}
