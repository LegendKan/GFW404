package controllers
import(
    "time"
    "ssmm/models"
    "fmt"
    "strconv"
)
func StartTimer() {
    go func() {
        for {
            erverdayFunc()
            now := time.Now()
            // 计算下一个零点
            next := now.Add(time.Hour * 24)
            next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
            t := time.NewTimer(next.Sub(now))
            <-t.C
        }
    }()
}

//产生账单，stop容器，发送邮件等
func erverdayFunc() {
    fmt.Println("Start the Timer")
    //获取全部active的accounts
    accounts, err := models.GetAllActiveAccountNew()
    if err!=nil{
        fmt.Println(err)
        return
    }
    now := time.Now()
    for _, v:=range accounts{
        duration:= v.Expiretime.Sub(now)
        days:= duration/time.Hour/24
        if days>5{
            continue
        }else if days>-3{
            //检查账单是否生成
            b, found :=models.GetUnpaidBillByAccount(v)
            if !found {
                //创建
                b.Createtime=now
                b.Expiretime=v.Expiretime
                b.Ispaid=0
                b.Price=v.Recurringprice
                var billingids string
                billingids = strconv.FormatInt(time.Now().Unix(), 10)
                b.Payno=billingids
                models.AddBill(&b)
            }
            //发邮件提醒

            user, erru:=models.GetUserById(v.Userid.Id)
            if erru==nil{
                go SendBillInfo(user.Email, user.Username, b.Payno, b.Price, b.Createtime.Format("2006-01-02"), b.Expiretime.Format("2006-01-02"))
            }

        }else if days>-5{
            //stop
            //获取server信息
            s, err:=models.GetServerById(v.Serverid.Id)
            if err!=nil {
                continue
            }
            StopContainer(s.Ip, s.Port, s.Auth, v.Containerid)
            //发邮件提醒

        }else{
            //delete并设置bill无效
            //获取server信息
            s, err:=models.GetServerById(v.Serverid.Id)
            if err!=nil {
                continue
            }
            DeleteContainer(s.Ip, s.Port, s.Auth, v.Containerid)
            b, found :=models.GetUnpaidBillByAccount(v)
            if found {
                b.Active=0
                e:=models.UpdateBillById(&b)
                if e!=nil {
                    fmt.Println(e)
                }
            }
            //发邮件提醒
            
        }
    }

    //新的购买没进行支付

}