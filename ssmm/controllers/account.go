package controllers

import (
	"encoding/json"
	"errors"
	"ssmm/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for Account
type AccountController struct {
	beego.Controller
}

func (c *AccountController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

//update
func (c *AccountController) Update() {
	id,_:=c.GetInt("id")
	containerid:=c.GetString("containerid")
	port,_:=c.GetInt("port")
	ip:=c.GetString("ip")
	password:=c.GetString("password")
	auth:=c.GetString("auth")

	var sortby []string
    var order []string
    var limit int64 = 10000
    var offset int64 = 0
	fields := []string{"Id", "Auth"}
    query := map[string]string{
        "Ip": ip,
    }
    server,err:=models.GetAllActiveServer(query, fields, sortby, order, offset, limit)
    if err!=nil{
    	c.Abort("500")
    }
    if len(server)==0{
    	c.Abort("404")
    }else if auth==((models.Server)(server[0])).Auth{
    	var v models.Account
    	v.Id=id
    	v.Containerid=containerid
    	v.Port=port
    	v.Password=password
    	models.UpdateAccountById(&v)
    	c.Data["json"]="OK"
    }else{
    	c.Abort("403")
    }
	
	c.ServeJson()
}

// @Title Post
// @Description create Account
// @Param	body		body 	models.Account	true		"body for Account content"
// @Success 200 {int} models.Account.Id
// @Failure 403 body is empty
// @router / [post]
func (c *AccountController) Post() {
	var v models.Account
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if id, err := models.AddAccount(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Account by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Account
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AccountController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAccountById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}


// @Title Get All
// @Description get Account
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Account
// @Failure 403
// @router / [get]
func (c *AccountController) GetAll() {
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

	l, err := models.GetAllAccount(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Account
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Account	true		"body for Account content"
// @Success 200 {object} models.Account
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AccountController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Account{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateAccountById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Account
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AccountController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAccount(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
