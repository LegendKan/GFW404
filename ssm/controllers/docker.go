package controllers

import (
	"encoding/json"
	"ssm/controllers/unixsocket"
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

func ListContainers() string {
	url := baseDocker + "/containers/json?all=1"
	_, result := unixsocket.UnixSocket("GET", url, "")
	return result
}

func DisplaySysInfo() string {
	statusCode, ret := unixsocket.UnixSocket("GET", baseDocker+"/info", "")
	if statusCode == 200 {
		return ret
	}
	return ""
}
