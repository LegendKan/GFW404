package controllers
import(
    "strconv"
)

type MsgRet struct {
    Status bool      `json:"status"`
    Msg    string	 `json:"data"`
}

//stop container
func StopContainer(ip string, port int, auth string, cid string) (MsgRet, error){
    var msg MsgRet
    resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/docker/stop?auth="+auth+"&cid="+cid)
    if err != nil {
        return msg, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
        return msg, err
    }
    fmt.Println("Stop Container:", string(body))
    err = json.Unmarshal(body, &msg)
    if err != nil {
        return msg, err
    }
    return msg, nil
}

//delete container
func DeleteContainer(ip string, port int, auth string, cid string) (MsgRet, error){
    var msg MsgRet
    resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/docker/delete?auth="+auth+"&cid="+cid)
    if err != nil {
        return msg, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
        return msg, err
    }
    fmt.Println("Stop Container:", string(body))
    err = json.Unmarshal(body, &msg)
    if err != nil {
        return msg, err
    }
    return msg, nil
}