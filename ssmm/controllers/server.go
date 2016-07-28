package controllers

import (
	"encoding/json"
	"errors"
	"ssmm/models"
	"strconv"
	"strings"
"fmt"
	"github.com/astaxie/beego"
)

// oprations for Server
type ServerController struct {
	beego.Controller
}

func (c *ServerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Server
// @Param	body		body 	models.Server	true		"body for Server content"
// @Success 200 {int} models.Server.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ServerController) Post() {
	var v models.Server
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	//查询该IP的服务器是否已经存在了，如果存在则验证密码，对则返回现在该服务器上的容器
	var sortby []string
    var order []string
    var limit int64 = 10000
    var offset int64 = 0
	fields := []string{"Id", "Auth"}
    query := map[string]string{
        "Ip": v.Ip,
    }
    server,err:=models.GetAllActiveServer(query, fields, sortby, order, offset, limit)
    if err!=nil{
    	c.Abort("500")
    }
    if len(server)==0{
    	if _, err := models.AddServer(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Abort("500")
		}
    }else if v.Auth==server[0].Auth{
    	//获取现在active的容器
    	fields1 := []string{"Id", "Containerid", "Port", "Password"}
	    query1 := map[string]string{
	        "Serverid": strconv.Itoa(server[0].Id),
	        "Active":"1",
	    }
	    accounts,err:=models.GetAllAccount(query1, fields1, sortby, order, offset, limit)
	    if err!=nil{
	    	c.Abort("500")
	    }

	    simples:=make([]models.SimpleAccount, 0)
	    for _, single:= range accounts{
	    	var tmp models.SimpleAccount
	    	tmp.Containerid=single.Containerid
	    	tmp.Password=single.Password
	    	tmp.Port=single.Port
	    	tmp.Id=single.Id
	    	tmp.Status=single.Active
	    	simples=append(simples,tmp)
	    }
	    fmt.Println(simples)
	    // r,_:=json.Marshal(simples)
	    // fmt.Println(r)
	    c.Data["json"]=simples
    }else{
    	//update server
    	v.Id=server[0].Id
    	v.Have=server[0].Have
    	models.UpdateServerById(&v)
    	c.Data["json"] = ""
    }
	
	c.ServeJson()
}

// @Title Get
// @Description get Server by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Server
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ServerController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetServerById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Server
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Server
// @Failure 403
// @router / [get]
func (c *ServerController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllServer(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Server
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Server	true		"body for Server content"
// @Success 200 {object} models.Server
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ServerController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Server{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateServerById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Server
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ServerController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteServer(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
