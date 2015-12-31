package controllers

import (
	"encoding/json"
	"ssm/controllers/unixsocket"
	"fmt"
	"ssm/models"
	"net/http"
)

var baseDocker string
var baseImage string

func init() {
	baseDocker = "unix:///var/run/docker.sock:"
	baseImage = "shadowsocks"
}

func AddContainer(port, password string) (string, bool) {
	data := "{\"Cmd\":[\"-k\", \"" + password + "\"],\"Image\":\"" + baseImage + "\", \"ExposedPorts\":{ \"" + innerPort + "/tcp\": {} }"
	if len(port) > 0 {
		data += ",\"HostConfig\": {\"PortBindings\":{ \"" + innerPort + "/tcp\": [{ \"HostPort\": \"" + port + "\" }] }} }"
	} else {
		data += ",\"HostConfig\": {\"PublishAllPorts\":true}}"
	}
	/*"\"Hostname\":\"\"," +
	"\"Domainname\": \"\"," +
	"\"User\":\"\"," +
	"\"Memory\":0," +
	"\"MemorySwap\":0," +
	"\"AttachStdin\":false," +
	"\"AttachStdout\":true," +
	"\"AttachStderr\":true," +
	"\"Tty\":false," +
	",\"HostConfig\": {\"PortBindings\":{ \"8000/tcp\": [{ \"HostPort\": \"" + port + "\" }] }} }"*/
	statusCode, result := unixsocket.UnixSocket("POST", baseDocker+"/containers/create", data)
	ret := make(map[string]interface{})
	err := json.Unmarshal([]byte(result), &ret)
	if err != nil {
		return "", false
	}
	id := ret["Id"].(string)
	if statusCode == 201 && StartContainer(id) {
		return id, true
	}
	return "", false
}

func StartContainer(id string) bool {
	statusCode, _ := unixsocket.UnixSocket("POST", baseDocker+"/containers/"+id+"/start", "")
	if statusCode == 204 || statusCode == 304 {
		return true
	}
	return false
}

func InspectContainer(id string) string {
	statusCode, ret := unixsocket.UnixSocket("GET", baseDocker+"/containers/"+id+"/json", "")
	if statusCode == 200 {
		return ret
	}
	return ""
}

func StopContainer(id string) bool {
	statusCode, _ := unixsocket.UnixSocket("POST", baseDocker+"/containers/"+id+"/stop", "")
	if statusCode == 204 || statusCode == 304 || statusCode == 404 {
		return true
	}
	return false
}

func RestartContainer(id string) bool {
	statusCode, _ := unixsocket.UnixSocket("POST", baseDocker+"/containers/"+id+"/restart", "")
	if statusCode == 204 {
		return true
	}
	return false
}

func KillContainer(id string) bool {
	statusCode, _ := unixsocket.UnixSocket("POST", baseDocker+"/containers/"+id+"/kill", "")
	if statusCode == 204 || statusCode == 404 {
		return true
	}
	return false
}

func DeleteContainer(id string) bool {
	statusCode, _ := unixsocket.UnixSocket("DELETE", baseDocker+"/containers/"+id+"?v=1&force=1", "")
	if statusCode == 204 || statusCode == 404 {
		return true
	}
	return false
}

func ListContainers() ([]models.Container, bool) {
	url := baseDocker + "/containers/json?all=1"
	statusCode, result := unixsocket.UnixSocket("GET", url, "")
	if statusCode ==200{
		fmt.Println(result)
		containers := make([]models.Container,0)
		if err:=json.Unmarshal([]byte(result), &containers);err!=nil{
			return nil, false
		}
		sscs:=make([]models.Container,len(containers))
		for _,v := range containers{
			if v.Image==baseImage{
				sscs=append(sscs, v)
			}
		}
		return sscs, true
	}
	return nil, false
}

func SyncContainers(accounts []models.Account) bool{
	containers:=ListContainers()
	var has bool
	for _, v:=range accounts{
		has=false
		for _, con:=range containers{
			if con.Id==v.Containerid{
				has=true
				if strings.Contains(con.Status, "created")||strings.Contains(con.Status, "paused")||strings.Contains(con.Status, "exited"){
					//start container
					if !StartContainer(con.Id){
						has=false
					}
				}
				break
			}
		}
		if !has{
			//创建新的容器，并发送到服务器
			resp, err := http.Get("http://example.com/")
		}
	}
	return false
}

func DisplaySysInfo() string {
	statusCode, ret := unixsocket.UnixSocket("GET", baseDocker+"/info", "")
	if statusCode == 200 {
		return ret
	}
	return ""
}
