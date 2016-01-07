package controllers

import (
	"errors"
	"ssmm/models"
	"strconv"
)

type WebUserController struct {
	baseController
}

func (c *WebUserController) GetRegister() {
	c.Data["IsUser"] = true
	c.TplNames = "register.html"
}

func (c *WebUserController) Register() {
	c.Data["IsUser"] = true
	var user models.User
	email := c.GetString("email")
	pass := c.GetString("password")
	//from:=c.Ctx.Request.Referer()
	if len(email) <= 8 || len(pass) <= 4 {
		c.Data["haserror"] = true
		c.Data["error"] = "请填写正确的邮箱、密码"
		c.TplNames = "register.html"
		//c.Render()
		return
	}
	//还有点问题
	u, _ := models.GetUserByEmail(email)
	if u != nil {
		c.Data["haserror"] = true
		c.Data["error"] = "邮箱已注册"
		c.TplNames = "register.html"
		return
	}
	firstname := c.GetString("firstname")
	lastname := c.GetString("lastname")
	username := firstname + lastname
	user.Email = email
	user.Username = username
	user.Password = pass
	user.Firstname=firstname
	user.Lastname=lastname
	user.Country=c.GetString("country")
	user.Province=c.GetString("state")
	user.City=c.GetString("city")
	user.Company=c.GetString("companyname")
	user.Address1=c.GetString("address1")
	user.Address2=c.GetString("address2")
	user.Zipcode=c.GetString("postcode")
	uid, err := models.AddUser(&user)
	if err != nil {
		c.Data["haserror"] = true
		c.Data["error"] = err.Error()
		c.TplNames = "register.html"
		return
	}
	go SendWelcome(email,username,pass)
	//添加session,cookie
	c.SetSession("email", email)
	c.SetSession("uid", int(uid))
	c.SetSession("username", username)
	c.Ctx.SetCookie("email", email)
	c.Ctx.SetCookie("password", pass)
	c.Redirect("/user", 302)
}

func (c *WebUserController) GetHome() {
	c.Data["IsUser"] = true
	email := c.GetSession("email")
	uid := c.GetSession("uid")
	userid,_:=uid.(int)
	if email == nil||uid==nil {
		c.Data["IsUser"] = true
		c.Redirect("/user/login", 302)
		return
	}
	//获取正常的服务GetAllAccount
	maps,err:=models.GetAllAccountByUserId(strconv.Itoa(userid))
	//fmt.Println(maps)
	if err==nil {
		c.Data["accounts"]=maps
	}
	//获取未付款的账单
	bills,err1:=models.GetAllUnpaidBills(strconv.Itoa(userid))
	if err1==nil {
		c.Data["bills"]=bills
	}
	c.Data["IsUser"] = true
	username := c.GetSession("username")
	name,_:=username.(string)
	c.Data["username"]=name
	c.TplNames = "userhome.html"
}

func (c *WebUserController) GetLogin() {
	c.Data["IsUser"] = true
	c.TplNames = "login.html"
}

func (c *WebUserController) Logout() {
	c.Data["IsUser"] = true
	c.DelSession("email")
	c.DelSession("uid")
	c.DelSession("username")
	c.Redirect("/user/login", 302)
}

func (c *WebUserController) Login() {
	c.Data["IsUser"] = true
	email := c.GetString("username")
	pass := c.GetString("password")
	if len(email) <= 8 || len(pass) <= 4 {
		c.Data["haserror"] = true
		c.Data["error"] = "请填写正确的邮箱、密码"
		c.TplNames = "login.html"
		return
	}
	// u, _ := models.GetUserByEmail(email)
	// if u == nil || u.Password != pass {
	// 	c.Data["haserror"] = true
	// 	c.Data["error"] = "邮箱或密码错误"
	// 	c.TplNames = "login.html"
	// 	return
	// }
	uid, username, err := verifyLogin(email, pass)
	if err != nil {
		c.Data["haserror"] = true
		c.Data["error"] = err.Error()
		c.TplNames = "login.html"
		return
	}
	c.SetSession("email", email)
	c.SetSession("uid", uid)
	c.SetSession("username", username)
	c.Ctx.SetCookie("email", email)
	c.Ctx.SetCookie("password", pass)
	c.Redirect("/user", 302)
}

func verifyLogin(email, pass string) (uid int, username string, err error) {
	u, _ := models.GetUserByEmail(email)
	if u == nil || u.Password != pass {
		//ret = false
		err = errors.New("邮箱或密码错误")
		return 0, "", err
	}
	return u.Id, u.Username, nil
}


func (c *WebUserController) GetDetail() {
	c.Data["IsUser"] = true
	email := c.GetSession("email")
	uid := c.GetSession("uid")
	userid,_:=uid.(int)
	if email == nil||uid==nil {
		c.Data["IsUser"] = true
		c.Redirect("/user/login", 302)
		return
	}

	idStr := c.Ctx.Input.Params[":id"]
	id, err1 := strconv.Atoi(idStr)
	if err1!=nil {
		c.Data["error"] = "参数出错了！"
		c.TplNames = "error.html"
		return
	}

	v, err := models.GetAccountDetailById(id)
	if err != nil {
		c.Data["error"] = err.Error()
		c.TplNames = "error.html"
		return
	} else {
		//验证用户信息
		tid,_:=strconv.Atoi(v["userid"].(string))
		if userid!=tid || v["active"]!="1" {
			c.Data["ip"] = "***.***.***.***"
			c.Data["port"] = 250
			c.Data["password"] = "U Guess"
		}else{
			c.Data["ip"] = v["ip"]
			c.Data["port"] = v["port"]
			c.Data["password"] = v["password"]
		}	
	}
	c.TplNames = "accountdetail.html"
}

func (c *WebUserController) PayBill(){
	c.Data["IsUser"] = true
	idStr := c.Ctx.Input.Params[":id"]
	id, err1 := strconv.Atoi(idStr)
	if err1!=nil{
		c.Data["error"] = "参数出错了！"
		c.TplNames = "error.html"
		return
	}
	b,err:=models.GetBillById(id)
	if err!=nil {
		c.Data["error"] = "出错了！"
		c.TplNames = "error.html"
		return
	}
	params:=CreatePay(b.Price,b.Payno,strconv.Itoa(b.Accountid)+"Shadowsocks账号")
	c.Data["total"] = b.Price
	c.Data["params"] = params
	c.TplNames = "paynow.html"
}
