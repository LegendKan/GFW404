package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"ssmm/models"
	"strconv"
	"strings"
	"time"
)

type CartController struct {
	baseController
}

type CartItem struct {
	Server *models.Server
	Cycle  int8
	Price  float64
	CartID int
}

func (c *CartController) ConfService() {
	id := c.Ctx.Input.Param(":id")
	b, error := strconv.Atoi(id)
	if error != nil {
		c.Abort("404")
	}
	server, err := models.GetServerById(b)
	if err != nil || server.Isonline == 0 || server.Remain <= 0 {
		c.Abort("404")
	}
	server.Description = strings.Replace(server.Description, "|", "<br />", -1)
	c.Data["service"] = server
	c.Data["IsService"] = true
	c.TplNames = "confproduct.html"
}

func (c *CartController) CheckoutService() {
	c.Data["IsService"] = true
	//this.Data["Title"] = "产品与服务"
	items := c.GetSession("cartitems")
	if items == nil {
		items = make([]CartItem, 0)
	}
	itemss := items.([]CartItem)
	c.Data["cartitems"] = itemss
	//fmt.Println(len(itemss), itemss[1].Server.Title)
	var totalprice float64
	for _, item := range itemss {
		totalprice += item.Price
	}
	c.Data["total"] = totalprice
	c.TplNames = "cartview-dev.html"
}

func (c *CartController) AddService() {
	c.Data["IsService"] = true
	//this.Data["Title"] = "产品与服务"
	serviceid, _ := c.GetInt("serviceid")
	cycle := c.GetString("billingcycle")
	server, err := models.GetServerById(serviceid)
	if err != nil || server.Isonline == 0 || server.Remain <= 0 {
		c.Abort("404")
	}
	var cycletype int8
	var price float64
	if cycle == "monthly" {
		cycletype = 1
		price = server.Month
	} else if cycle == "quarterly" {
		cycletype = 2
		price = server.Quarter
	} else {
		cycletype = 3
		price = server.Year
	}
	items := c.GetSession("cartitems")
	if items == nil {
		items = make([]CartItem, 0)
	}
	itemss := items.([]CartItem)
	count := len(itemss)
	item := CartItem{server, cycletype, price, count}
	itemss = append(itemss, item)
	c.SetSession("cartitems", itemss)
	fmt.Println(serviceid, cycle, item.Server.Title)
	c.Redirect("/cart/view", 302)
}

func (c *CartController) ViewService() {
	c.Data["IsService"] = true
	//this.Data["Title"] = "产品与服务"
	c.TplNames = "cart.html"
}

func (c *CartController) PlaceOrder() {
	items := c.GetSession("cartitems")
	if items == nil {
		c.Redirect("/service", 302)
		return
	}
	itemss := items.([]CartItem)
	if len(itemss) <= 0 {
		c.Redirect("/service", 302)
		return
	}
	user := c.GetSession("uid")
	var userid int
	if user == nil {
		//判断是登录还是注册
		isnew, err := c.GetBool("newuser")
		if err != nil {
			c.Data["haserror"] = true
			c.Data["error"] = err.Error()
			c.TplNames = "cartview-dev.html"
		}
		if isnew {
			//注册
			var user models.User
			email := c.GetString("email1")
			pass := c.GetString("password1")
			if len(email) <= 8 || len(pass) <= 4 {
				c.Data["haserror"] = true
				c.Data["error"] = "请填写正确的邮箱、密码"
				c.TplNames = "cartview-dev.html"
				//c.Render()
				return
			}
			//还有点问题
			u, _ := models.GetUserByEmail(email)
			if u != nil {
				c.Data["haserror"] = true
				c.Data["error"] = "邮箱已注册"
				c.TplNames = "cartview-dev.html"
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
				c.TplNames = "cartview-dev.html"
				return
			}
			//添加session,cookie
			c.SetSession("email", email)
			c.SetSession("uid", uid)
			userid = int(uid)
			c.SetSession("username", username)
			c.Ctx.SetCookie("email", email)
			c.Ctx.SetCookie("password", pass)
		} else {
			//登录
			email := c.GetString("username")
			pass := c.GetString("password")
			if len(email) <= 8 || len(pass) <= 4 {
				c.Data["haserror"] = true
				c.Data["error"] = "请填写正确的邮箱、密码"
				c.TplNames = "cartview-dev.html"
				return
			}
			uid, username, err := verifyLogin(email, pass)
			if err != nil {
				c.Data["haserror"] = true
				c.Data["error"] = err.Error()
				c.TplNames = "cartview-dev.html"
				return
			}
			c.SetSession("email", email)
			c.SetSession("uid", uid)
			userid = uid
			c.SetSession("username", username)
			c.Ctx.SetCookie("email", email)
			c.Ctx.SetCookie("password", pass)
		}
	} else {
		userid = user.(int)
	}
	//要付款总价
	var total float64
	var billingids string
	billingids = strconv.FormatInt(time.Now().Unix(), 10)
	timenow := time.Now()
	var exprietime time.Time
	for _, item := range itemss {
		//创建服务和账单
		if item.Cycle == 1 {
			exprietime = timenow.AddDate(0, 1, 0)
		} else if item.Cycle == 2 {
			exprietime = timenow.AddDate(0, 3, 0)
		} else {
			exprietime = timenow.AddDate(1, 0, 0)
		}
		account := &models.Account{Serverid: item.Server, Userid: &models.User{Id: userid}, Cycle: item.Cycle, Expiretime: exprietime}
		aid, err := models.AddAccount(account)
		if err != nil {
			c.Data["haserror"] = true
			c.Data["error"] = err.Error()
			c.TplNames = "cartview-dev.html"
			return
		}
		billing := &models.Bill{Accountid: int(aid), Price: item.Price, Payno: billingids, Expiretime: timenow.AddDate(0, 0, 5)}
		_, err = models.AddBill(billing)
		//billingid := int(bid)
		if err != nil {
			c.Data["haserror"] = true
			c.Data["error"] = err.Error()
			c.TplNames = "cartview-dev.html"
			return
		}
		total += item.Price

	}
	c.SetSession("cartitems", nil)
	//c.Redirect("/cart/view", 302)
	c.Data["total"] = total
	flash := beego.NewFlash()
	//flash.Set("billing", billingids)
	flash.Data["billing"] = billingids
	flash.Data["total"] = strconv.FormatFloat(total, 'e', 2, 32)
	flash.Store(&c.Controller)
	//flash.Set("total", total)
	c.TplNames = "payonline.html"
}
