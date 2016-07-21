package controllers

import (
	"encoding/json"
	"ssm/controllers/unixsocket"
	"fmt"
	"ssm/models"
	"net/http"
	"strings"
	"net/url"
	"github.com/astaxie/beego"
	"strconv"
)

var baseSocket string

func init() {
	baseSocket = "unix:///var/run/shadowsocks.sock:"
}

func addPort(port, pass string) bool {
	data := "add: {\"server_port\": "+port+", \"password\":\""+pass+"\"}"
	return true
}

func stopPort(port string) bool {
	
	return true
}

func deletePort(port string) bool {
	return stopPort(port)
}