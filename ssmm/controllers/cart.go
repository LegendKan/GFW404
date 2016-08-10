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
	RecurringPrice float64
	CartID int
	Password string
}

func (c *CartController) ConfService() {
	id := c.Ctx.Input.Param(":id")
	b, error := strconv.Atoi(id)
	if error != nil {
		c.Abort("404")
	}
	server, err := models.GetServerById(b)
	if err != nil || server.Isonline == 0 || server.Amount-server.Have <= 0 {
		c.Abort("404")
	}
	server.Description = strings.Replace(server.Description, "|", "<br />", -1)
	c.Data["service"] = server
	c.Data["IsService"] = true
	c.TplName = "confproduct.html"
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
	//origin ip
	//ipindex:=strings.Index(c.Ctx.Request.RemoteAddr,":")
	//c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
	//get real ip from nginx
	c.Data["clientip"]=c.Ctx.Request.Header.Get("X-Real-IP")
	c.TplName = "cartview-dev.html"
}

func (c *CartController) AddService() {
	c.Data["IsService"] = true
	//this.Data["Title"] = "产品与服务"
	serviceid, _ := c.GetInt("serviceid")
	cycle := c.GetString("billingcycle")
	password := c.GetString("password")
	server, err := models.GetServerById(serviceid)
	if err != nil || server.Isonline == 0 || server.Amount-server.Have <= 0 {
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
	item := CartItem{server, cycletype, price, price, count, password}
	itemss = append(itemss, item)
	c.SetSession("cartitems", itemss)
	fmt.Println(serviceid, cycle, item.Server.Title)
	c.Redirect("/cart/view", 302)
}

func (c *CartController) DeleteService() {
	c.Data["IsService"] = true
	id := c.Ctx.Input.Param(":id")
	b, error := strconv.Atoi(id)
	if error != nil {
		c.Redirect("/cart/view", 302)
		return
	}
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
	if len(itemss)-1<b {
		c.Redirect("/cart/view", 302)
		return
	}
	itemss = append(itemss[:b], itemss[b+1:]...)
	c.SetSession("cartitems", itemss)
	// var totalprice float64
	// for _, item := range itemss {
	// 	totalprice += item.Price
	// }
	// ipindex:=strings.Index(c.Ctx.Request.RemoteAddr,":")
	// c.Data["cartitems"] = itemss
	// c.Data["total"] = totalprice
	// c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
	// c.TplName = "cartview-dev.html"
	c.Redirect("/cart/view", 302)
}

func (c *CartController) ClearService() {
	c.Data["IsService"] = true
	c.DelSession("cartitems")
	c.Redirect("/cart/view", 302)
}

func (c *CartController) ViewService() {
	c.Data["IsService"] = true
	//this.Data["Title"] = "产品与服务"
	c.TplName = "cart.html"
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func calTotalPrice(itemss []CartItem) float64 {
	totalprice float64
	for _, item := range itemss {
		totalprice += item.Price
	}
	return totalprice
}

func (c *CartController) PromoteFilter() {
	c.Data["IsService"] = true
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
	
	couponCode := c.GetString("couponcode")
	if len(couponCode)<=0{
		c.Redirect("/service", 302)
		return
	}
	var totalprice, recurringprice float64
	coupon, _ := models.GetCouponByCode(couponCode)
	if coupon == nil{
		c.Data["haserror"] = true
		c.Data["error"] = "所填写优惠码不存在！"
		for _, item := range itemss {
			totalprice += item.Price
		}
		recurringprice = totalprice
	}else if coupon.Type != -1 && coupon.Usedtimes >= coupon.Totaltimes{
		c.Data["haserror"] = true
		c.Data["error"] = "优惠码使用达到上限！"
		totalprice = calTotalPrice(itemss)
		recurringprice = totalprice
	}else if now := time.Now(); now.Before(coupon.Effecttime) || now.After(coupon.Expiretime){
		c.Data["haserror"] = true
		c.Data["error"] = "优惠码已过期或者尚未生效！"
		totalprice = calTotalPrice(itemss)
		recurringprice = totalprice
	}else{
		serverids = strings.Split(coupon.Serverids,"|")
		for _, item := range itemss {
			if (coupon.Serverids == "*" || stringInSlice(strconv.Itoa(item.Server.Id),serverids)) && (coupon.Cycle == 3 || item.Cycle == coupon.Cycle){
				oldvalue := item.Price
				if coupon.Type == 0 {
					item.Price = item.Price * coupon.Content/100
				}else if coupon.Type == 1{
					item.Price -= coupon.Content
				}else if coupon.Type ==2 {
					item.Price = coupon.Content
				}else{
					item.Price = item.Price
				}

				if item.Price <=0 {
					item.Price = oldvalue
				}

				if coupon.Recursion ==1{
					item.RecurringPrice = item.Price
				}
			}
			totalprice += item.Price
			recurringprice += item.RecurringPrice
		}
	}
	
	ipindex:=strings.Index(c.Ctx.Request.RemoteAddr,":")
	c.Data["cartitems"] = itemss
	c.Data["total"] = totalprice
	c.Data["recurring"] = recurringprice
	c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
	c.TplName = "cartview-dev.html"
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
	var totalprice float64
	for _, item := range itemss {
		totalprice += item.Price
	}
	var username, email string
	ipindex:=strings.Index(c.Ctx.Request.RemoteAddr,":")
	if user == nil {
		//判断是登录还是注册
		isnew, err := c.GetBool("newuser")
		if err != nil {
			c.Data["haserror"] = true
			c.Data["error"] = err.Error()
			c.Data["cartitems"] = itemss
			c.Data["total"] = totalprice
			c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
			c.TplName = "cartview-dev.html"
			return
		}
		if isnew {
			//注册
			var user models.User
			email = c.GetString("email1")
			pass := c.GetString("password1")
			if len(email) <= 8 || len(pass) <= 4 {
				c.Data["haserror"] = true
				c.Data["error"] = "请填写正确的邮箱、密码"
				c.Data["cartitems"] = itemss
				c.Data["total"] = totalprice
				c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
				c.TplName = "cartview-dev.html"
				//c.Render()
				return
			}
			//还有点问题
			u, _ := models.GetUserByEmail(email)
			if u != nil {
				c.Data["haserror"] = true
				c.Data["error"] = "邮箱已注册"
				c.Data["cartitems"] = itemss
				c.Data["total"] = totalprice
				c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
				c.TplName = "cartview-dev.html"
				return
			}
			firstname := c.GetString("firstname")
			lastname := c.GetString("lastname")
			username = firstname + lastname
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
				c.Data["cartitems"] = itemss
				c.Data["total"] = totalprice
				c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
				c.TplName = "cartview-dev.html"
				return
			}
			//发邮件
			go SendWelcome(email,username,pass)
			//添加session,cookie
			c.SetSession("email", email)
			c.SetSession("uid", uid)
			userid = int(uid)
			c.SetSession("username", username)
			c.Ctx.SetCookie("email", email)
			c.Ctx.SetCookie("password", pass)
		} else {
			//登录
			email = c.GetString("username")
			pass := c.GetString("password")
			if len(email) <= 8 || len(pass) <= 4 {
				c.Data["haserror"] = true
				c.Data["error"] = "请填写正确的邮箱、密码"
				c.Data["cartitems"] = itemss
				c.Data["total"] = totalprice
				c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
				c.TplName = "cartview-dev.html"
				return
			}
			uid, usern, err := verifyLogin(email, pass)
			username=usern
			if err != nil {
				c.Data["haserror"] = true
				c.Data["error"] = err.Error()
				c.Data["cartitems"] = itemss
				c.Data["total"] = totalprice
				c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
				c.TplName = "cartview-dev.html"
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
		emailtmp := c.GetSession("email")
		usertmp := c.GetSession("username")
		email,_=emailtmp.(string)
		username,_=usertmp.(string)
	}
	//要付款总价
	var total float64
	var billingids string
	billingids = strconv.FormatInt(time.Now().Unix(), 10)
	timenow := time.Now()
	var exprietime time.Time
	var cyclestr string
	for _, item := range itemss {
		//创建服务和账单
		if item.Cycle == 1 {
			exprietime = timenow.AddDate(0, 1, 0)
			cyclestr="月付"
		} else if item.Cycle == 2 {
			exprietime = timenow.AddDate(0, 3, 0)
			cyclestr="季付"
		} else {
			exprietime = timenow.AddDate(1, 0, 0)
			cyclestr="年付"
		}
		account := &models.Account{Serverid: item.Server, Password: item.Password, Userid: &models.User{Id: userid}, Cycle: item.Cycle, Expiretime: exprietime, Firstprice: item.Price, Recurringprice: item.Price}
		aid, err := models.AddAccount(account)
		if err != nil {
			c.Data["haserror"] = true
			c.Data["error"] = err.Error()
			c.Data["cartitems"] = itemss
			c.Data["total"] = totalprice
			c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
			c.TplName = "cartview-dev.html"
			return
		}
		billing := &models.Bill{Accountid: int(aid), Price: item.Price, Payno: billingids, Expiretime: timenow.AddDate(0, 0, 5)}
		_, err = models.AddBill(billing)
		//billingid := int(bid)
		if err != nil {
			c.Data["haserror"] = true
			c.Data["error"] = err.Error()
			c.Data["cartitems"] = itemss
			c.Data["total"] = totalprice
			c.Data["clientip"]=c.Ctx.Request.RemoteAddr[0:ipindex]
			c.TplName = "cartview-dev.html"
			return
		}
		total += item.Price

	}
	//发邮件
	now :=time.Now()
	go SendBillInfo(email,username,billingids,total,now.Format("2006-01-02"),now.AddDate(0, 0, 5).Format("2006-01-02"))

	c.SetSession("cartitems", nil)
	//c.Redirect("/cart/view", 302)
	c.Data["total"] = total
	flash := beego.NewFlash()
	//flash.Set("billing", billingids)
	flash.Data["billing"] = billingids
	flash.Data["total"] = strconv.FormatFloat(total, 'e', 2, 32)
	flash.Store(&c.Controller)
	//flash.Set("total", total)
	//old for pingpp
	//c.TplName = "payonline.html"
	fmt.Println("I am here Catch me")
	//ipaynow.cn paynow支付
	params:=CreatePay(total,billingids,cyclestr+"Shadowsocks账号")
	c.Data["params"] = params
	c.TplName = "paynow.html"
	
}
