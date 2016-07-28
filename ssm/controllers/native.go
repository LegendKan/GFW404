package controllers

import (
	"fmt"
	"ssm/models"
	"strings"
	"net"
	"github.com/astaxie/beego"
	"os"
)

var baseSocket string
var conn *net.UnixConn
var buf [512]byte

func init() {
	socketype := "unixgram"
	os.Remove("/tmp/client.sock")
	laddr := net.UnixAddr{"/tmp/client.sock", socketype}
	var err error
	conn, err = net.DialUnix(socketype, &laddr/*can be nil*/,
	    &net.UnixAddr{"/var/run/shadowsocks.sock", socketype})
	if err != nil {
	    panic(err)
	}   
}

func recv() (string, error) {
	n, err := conn.Read(buf[0:])
	if err!=nil{
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println("recv "+)
	return string(buf[0:n]), nil
}

func addPort(port, pass string) (bool, err) {
	fmt.Println("add port "+port)
	data := "add: {\"server_port\": "+port+", \"password\":\""+pass+"\"}"
	_, err = conn.Write([]byte(data))
	if err != nil {
	    return false, err
	} 
	ret, err := recv()
	if err!=nil || ret!="ok"{
		return false, errors.New("return is not ok")
	}
	return true, nil
}

func stopPort(port string) (bool, err) {
	fmt.Println("remove port "+port)
	data := "remove: {\"server_port\":"+port+"}"
	_, err = conn.Write([]byte(data))
	if err != nil {
	    return false, err
	} 
	ret, err := recv()
	if err!=nil || ret!="ok"{
		return false, errors.New("return is not ok")
	}
	return true, nil
}

func deletePort(port string) bool {
	return stopPort(port)
}

func pintPort(port string) {
	fmt.Println("Ping...)
	data := "ping"
	_, err := conn.Write([]byte(data))
	if err != nil {
	    return false, err
	} 
	ret, err := recv()
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(ret)
}

func SyncPorts(accounts []models.Account) bool{
	fmt.Println("Start to Sync")
	for _, v:=range accounts{
		if v.Port>=MaxPort{
			MaxPort=v.Port+1
		}
		fmt.Println(MaxPort)
		//根据状态该删的删掉，该添加的添加
		if v.Status == 1 {
			addPort(strconv.Itoa(v.Port), v.Password)
		} else if v.Status == 2 {
			stopPort(strconv.Itoa(v.Port))
		} else if v.Status == 3 {
			deletePort(strconv.Itoa(v.Port))
		} else {
			fmt.Println("Unknown Status")
		}
	}
}