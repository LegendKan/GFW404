package controllers
import (
    "encoding/json"
	"crypto/md5"
	"encoding/hex"
    "github.com/astaxie/beego"
    "strconv"
    "time"
    "sort"
    "fmt"
    "ssmm/models"
    "io/ioutil"
    "net/http"
)

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

type PayNowController struct {
    baseController
}

func CreatePay(amt float64, orderid string, des string) map[string]string {
    params:=make(map[string]string)
    params["mhtOrderName"]="gfw404.com Shadowsocks账号"
    orderamt:=int(amt*100)
    params["mhtOrderAmt"]=strconv.Itoa(orderamt)
    params["mhtOrderDetail"]=des
    params["funcode"]="WP001"
    params["appId"]=beego.AppConfig.String("paynow::appid")
    params["mhtOrderNo"]=orderid
    params["mhtOrderType"]="01"
    params["mhtCurrencyType"]="156"
    params["mhtOrderTimeOut"]="3600"
    params["mhtOrderStartTime"]=time.Now().UTC().Add(8 * time.Hour).Format("20060102150405")
    params["notifyUrl"]=beego.AppConfig.String("paynow::notifyurl")
    params["frontNotifyUrl"]=beego.AppConfig.String("paynow::fronturl")
    params["mhtCharset"]="UTF-8"
    params["deviceType"]="02"
    params["mhtReserved"]="Welcome"
    params["mhtSignature"]=Signature(params,beego.AppConfig.String("paynow::securekey"))
    params["mhtSignType"]="MD5"

    return params
}

func Signature(params map[string]string, secret string) string {
	linked:=createLinkString(params)
	h := md5.New()
	h.Write([]byte(secret))
	sign:=hex.EncodeToString(h.Sum(nil))
	linked=linked+"&"+sign;
	h2:= md5.New()
	h2.Write([]byte(linked))
	return hex.EncodeToString(h2.Sum(nil))
}


func createLinkString(params map[string]string)string {
	
	sorted_keys := make([]string, 0)  
    for k, _ := range params {  
        sorted_keys = append(sorted_keys, k)  
    }  
    // sort 'string' key in increasing order  
    sort.Strings(sorted_keys)
    var tmp string
    flag:=1
    for _,k:=range sorted_keys{
        if k=="funcode" || k=="deviceType" || k=="mhtSignType" || k=="mhtSignature"{
            continue
        }
        if flag==1{
            tmp=k+"="+params[k]
            flag=0
        }else{
            tmp=tmp+"&"+k+"="+params[k]
        }
    }
    return tmp
    /*
    parameters := url.Values{}
    for k, v := range params {  
    	if k=="funcode" || k=="deviceType" || k=="mhtSignType" || k=="mhtSignature"{
    		continue
    	}
        parameters.Add(k, v)
    }
    return parameters.Encode()
    */
}

func (c *PayNowController) PayResult() {
    params:=make(map[string]string)
    params["funcode"]=c.GetString("funcode")
    params["appId"]=c.GetString("appId")
    params["mhtOrderNo"]=c.GetString("mhtOrderNo")
    params["mhtCharset"]=c.GetString("mhtCharset")
    params["tradeStatus"]=c.GetString("tradeStatus")
    params["mhtReserved"]=c.GetString("mhtReserved")
    params["signType"]=c.GetString("signType")
    params["signature"]=c.GetString("signature")
    result:=Verify(params,beego.AppConfig.String("paynow::securekey"))
    var payresult string
    if result{
        if params["tradeStatus"]=="A001"{
            payresult="恭喜您，支付成功"
        }else if params["tradeStatus"]=="A002"{
            payresult="很遗憾，支付失败"
        }else if params["tradeStatus"]=="A003"{
            payresult="支付结果未知"
        }else{
            payresult="状态未知"
        }
    }else{
        payresult="验签失败，不要搞破坏哟"
    }

    c.Data["payresult"] = payresult
    c.TplNames = "payresult.html"
}

func Verify(params map[string]string, secret string) bool {
    linked:=createLinkStringforVerify(params)
    h := md5.New()
    h.Write([]byte(secret))
    sign:=hex.EncodeToString(h.Sum(nil))
    linked=linked+"&"+sign;
    h2:= md5.New()
    h2.Write([]byte(linked))
    target:=hex.EncodeToString(h2.Sum(nil))
    if params["signature"]!=""&&params["signature"]==target{
        return true
    }else{
        return false
    }
}

func createLinkStringforVerify(params map[string]string)string{
    sorted_keys := make([]string, 0)  
    for k, _ := range params {  
        sorted_keys = append(sorted_keys, k)  
    }  
    // sort 'string' key in increasing order  
    sort.Strings(sorted_keys)
    var tmp string
    flag:=1
    for _,k:=range sorted_keys{
        if k=="signType" || k=="signature" || params[k]==""{
            continue
        }
        if flag==1{
            tmp=k+"="+params[k]
            flag=0
        }else{
            tmp=tmp+"&"+k+"="+params[k]
        }
    }
    return tmp
}

func (c *PayNowController) Callback() {
    params:=make(map[string]string)
    params["funcode"]=c.GetString("funcode")
    params["appId"]=c.GetString("appId")
    params["mhtOrderNo"]=c.GetString("mhtOrderNo")
    params["mhtOrderType"]=c.GetString("mhtOrderType")
    params["mhtCurrencyType"]=c.GetString("mhtCurrencyType")
    params["mhtOrderAmt"]=c.GetString("mhtOrderAmt")
    params["mhtOrderTimeOut"]=c.GetString("mhtOrderTimeOut")
    params["mhtOrderStartTime"]=c.GetString("mhtOrderStartTime")
    params["mhtCharset"]=c.GetString("mhtCharset")
    params["deviceType"]=c.GetString("deviceType")
    params["payChannelType"]=c.GetString("payChannelType")
    params["nowPayAccNo"]=c.GetString("nowPayAccNo")
    params["tradeStatus"]=c.GetString("tradeStatus")
    params["mhtReserved"]=c.GetString("mhtReserved")
    params["signType"]=c.GetString("signType")
    params["signature"]=c.GetString("signature")
    result:=Verify(params,beego.AppConfig.String("paynow::securekey"))
    if !result||params["tradeStatus"]!="A001"{
        c.Ctx.WriteString("success=Y")
        return
    }

    orderId := params["mhtOrderNo"]
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
        c.Ctx.WriteString("success=N")
        return
    }
    for _, bill := range l {
        //bill := billi.(*models.Bill)
        if bill.Ispaid == 1 {
            c.Ctx.WriteString("success=Y")
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
                        fmt.Println(orderId+"Create Container Error: ", err)
                    } else {
                        account.Active = 1
                        account.Containerid = con.Con.Containerid
                        account.Port, _ = strconv.Atoi(con.Con.Port)
                        account.Password = con.Con.Password
                        //models.UpdateAccountById(account)
                        server.Have--
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
    c.Ctx.WriteString("success=Y")
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

