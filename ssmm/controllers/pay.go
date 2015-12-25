package controllers

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	"log"
	"io/ioutil"
	"net/http"
	"ssmm/models"
	"strconv"
)

type NotifyCharge struct {
	Id              string                 `json:"id"`
	Object          string                 `json:"object"`
	Created         uint64                 `json:"created"`
	Livemode        bool                   `json:"livemode"`
	Paid            bool                   `json:"paid"`
	Refunded        bool                   `json:"refunded"`
	Order_no        string                 `json:"order_no"`
	App             string                 `json:"app"`
	Channel         string                 `json:"channel"`
	Amount          uint64                 `json:"amount"`
	Amount_settle   uint64                 `json:"amount_settle"`
	Amount_refunded uint64                 `json:"amount_refunded"`
	Time_expire     uint64                 `json:"time_expire"`
	Time_settle     uint64                 `json:"time_settle"`
	Transaction_no  string                 `json:"transaction_no"`
	Currency        string                 `json:"currency"`
	Client_ip       string                 `json:"client_ip"`
	Subject         string                 `json:"subject"`
	Body            string                 `json:"body"`
	Failure_code    int                    `json:"failure_code"`
	Failure_msg     string                 `json:"failure_msg"`
	Metadata        map[string]interface{} `json:"metadata"`
	//Refunds         RefundList             `json:"refunds"`
	//Credential      Credential             `json:"credential"`
}

type Container struct {
	Server      string `json:"server"`
	Password    string `json:"password"`
	Containerid string `json:"containerid"`
	Port        string `json:"port"`
}

type RetContainer struct {
	Status bool      `json:"status"`
	Con    Container `json:"data"`
}

type PayController struct {
	baseController
}

//new version pingpp
func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 2
	//设置 API Key
	pingpp.Key = "sk_test_jfbXD480KubTSOi1O0OW5SSC"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"
}

func (c *PayController) Pay() {
	flash := beego.ReadFromRequest(&c.Controller)
	billing := flash.Data["billing"]
	fmt.Println("billing", billing)
	amount := flash.Data["total"]
	if len(billing) <= 0 || len(amount) <= 0 {
		c.Redirect("/user", 302)
		return
	}
	total, _ := strconv.ParseFloat(amount, 64)
	total = 100 * total

	//new version pingpp
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	//extra 参数根据渠道不同有区别，下面注释的是一部分的示例
	extra := make(map[string]interface{})
	//upacp_wap
	extra["result_url"] = "http://www.gfw404.com:8080/pay/result/"
	params := &pingpp.ChargeParams{
		Order_no:  billing,
		App:       pingpp.App{Id: "app_vPu5yTeLaDa1jfT8"},
		Amount:    uint64(total),
		Channel:   "upacp_pc",
		Currency:  "cny",
		Client_ip: "127.0.0.1",
		Subject:   "ShadowSocks 账号",
		Body:      "ShadowSocks 账号 付款",
		Extra:     extra,
		Metadata:  metadata,
	}
	//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	ch, err := charge.New(params)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(ch)
	c.Data["json"] = ch
	c.ServeJson()
	/**old version
	// var jsonstring string
	// jsonstring = `{"order_no": "123456789011122233", "extra":{"result_url":"http://example.com/example/"}, "amount": 1,"app": {"id":"app_vPu5yTeLaDa1jfT8"},"channel": "upmp_wap","currency": "cny","client_ip": "192.168.1.1","subject": "test-subject","body": "test-body","metadata": {"color": "red"}}`
	var chargeParams pingpp.ChargeParams
	// json.Unmarshal([]byte(jsonstring), &chargeParams)
	chargeParams.Order_no = billing
	chargeParams.Extra.Result_url = "http://104.149.223.162:8080/pay/result/"
	chargeParams.App.Id = "app_vPu5yTeLaDa1jfT8"
	chargeParams.Channel = "upmp_wap"
	chargeParams.Currency = "cny"
	chargeParams.Client_ip = "192.168.1.1"
	chargeParams.Subject = "ShadowSocks 账号"
	chargeParams.Body = "ShadowSocks 账号 付款"
	chargeParams.Amount = uint64(total)
	meta := make(map[string]interface{})
	meta["color"] = "red"
	chargeParams.Metadata = meta
	chargeClient := pingpp.GetChargeClient("sk_test_jfbXD480KubTSOi1O0OW5SSC")
	charge, _ := chargeClient.New(&chargeParams)
	fmt.Printf("%v", charge)
	c.Data["json"] = charge
	c.ServeJson()
	*/
}

func (c *PayController) BeforePay() {
	c.TplNames = "pinus.html"
}

func (c *PayController) PayResult() {
	id := c.Ctx.Input.Param(":id")
	c.Data["payresult"] = id
	c.TplNames = "payresult.html"
}

func (c *PayController) Callback() {
	result := c.Ctx.Input.RequestBody
	fmt.Println("Callback:", string(result))
	//object := parseNotify(string(result))
	var charge NotifyCharge
	err0 := json.Unmarshal(result, &charge)
	if err0 != nil {
		fmt.Println(err0)
		c.Ctx.WriteString("success")
		return
	}
	//charge := object.(pingpp.NotifyCharge)
	orderId1 := charge.Order_no
	fmt.Println("OrderID:" + orderId1)
	//截取orderid
	//rs := []rune(orderId1)
	//orderId := string(rs[len(orderId1)-10:])
	orderId := orderId1[len(orderId1)-10 : len(orderId1)]
	//orderId := "1422697835"
	fmt.Println("OrderID:" + orderId)
	//通过order获得账单，从账单中获得Account信息，决定是创建还是延时
	var sortby []string
	var order []string
	var limit int64 = 100
	var offset int64 = 0
	fields := []string{"Id", "Accountid", "Ispaid"}
	query := map[string]string{
		"Payno": orderId,
	}
	l, err1 := models.GetAllBill(query, fields, sortby, order, offset, limit)
	if err1 != nil {
		c.Ctx.WriteString("success")
		return
	}
	for _, bill := range l {
		//bill := billi.(*models.Bill)
		if bill.Ispaid == 1 {
			c.Ctx.WriteString("success")
			return
		}
		bill.Ispaid = 1
		models.UpdateBillById(&bill)
		if account, err := models.GetAccountById(bill.Accountid); err == nil {
			//判断account是否已经创建
			if account.Active == 0 {
				//去创建
				serverid := account.Serverid
				fmt.Println(serverid.Description)
				//var aport int
				//var password string
				//serverid是完整的server对象还是只有ID的东西呢？
				if server, err := models.GetServerById(serverid.Id); err == nil {
					//获取创建必要的信息IP,port, pass
					ip := server.Ip
					port := server.Port
					pass := server.Auth
					//发起get(post)请求去创建
					con, err := createContainer(ip, port, pass)
					if err != nil {
						fmt.Println("Create Container Error: ", err)
					} else {
						account.Active = 1
						account.Containerid = con.Con.Containerid
						account.Port, _ = strconv.Atoi(con.Con.Port)
						account.Password = con.Con.Password
						//models.UpdateAccountById(account)
						server.Have++
						models.UpdateServerById(server)
						fmt.Println("ContainerID: ", con.Con.Containerid)
					}

					fmt.Println("Create Container: ", ip)
				}

			} else if account.Active == 1 {
				//延长时间
				timenow := account.Expiretime
				if account.Cycle == 1 {
					timenow.AddDate(0, 1, 0)
				} else if account.Cycle == 2 {
					timenow.AddDate(0, 3, 0)
				} else {
					timenow.AddDate(1, 0, 0)
				}
				account.Expiretime = timenow
			}
			//更新Account
			models.UpdateAccountById(account)
		}
	}
	c.Ctx.WriteString("success")
}

func createContainer(ip string, port int, auth string) (RetContainer, error) {
	var con RetContainer
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/add")
	if err != nil {
		return con, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return con, err
	}
	fmt.Println("Raw Container:", string(body))
	err = json.Unmarshal(body, &con)
	if err != nil {
		return con, err
	}
	return con, nil
}
