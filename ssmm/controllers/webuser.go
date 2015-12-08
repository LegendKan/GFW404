package controllers

import (
	"errors"
	"ssmm/models"
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
	uid, err := models.AddUser(&user)
	if err != nil {
		c.Data["haserror"] = true
		c.Data["error"] = err.Error()
		c.TplNames = "register.html"
		return
	}
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
	if email == nil {
		c.Data["IsUser"] = true
		c.Redirect("/user/login", 302)
		return
	}
	c.Data["IsUser"] = true
	c.TplNames = "userhome.html"
}

func (c *WebUserController) GetLogin() {
	c.Data["IsUser"] = true
	c.TplNames = "login.html"
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
