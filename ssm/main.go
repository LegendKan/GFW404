package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"net/http"
	"ssm/models"
	_ "ssm/routers"
)

func main() {
	/*"errors"
	var myhook = func() error {
		if contactMaster() {
			return nil
		}
		return errors.New("Test")
	}
	*/
	//beego.AddAPPStartHook(myhook)
	beego.Run()
}

func contactMaster() bool {
	masterport, _ := beego.AppConfig.Int("httpport")
	amount, _ := beego.AppConfig.Int("amount")
	have, _ := beego.AppConfig.Int("have")
	var server = models.Server{beego.AppConfig.String("server"), masterport, beego.AppConfig.String("passauth"), beego.AppConfig.String("location"),
		beego.AppConfig.String("description"), amount, have}
	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("json err:", err)
	}

	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post("http://"+beego.AppConfig.String("masteraddr")+":"+beego.AppConfig.String("masterport")+"/api/server", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if res.StatusCode == 200 {
		//处理已经创建的，进行恢复

		return true
	}

	return false
}
