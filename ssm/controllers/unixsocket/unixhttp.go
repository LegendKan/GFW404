package unixsocket

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func checkURL(urlstring string) (*url.URL, error) {
	u, err := url.Parse(urlstring)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "unix" {
		return nil, fmt.Errorf("Scheme must be unix ie. unix:///var/run/daemon/sock:/path")
	}
	return u, nil
}

func RequestSocket(method, urlstring, data, header, cookie string) (int, string) {
	u, err := checkURL(urlstring)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hostAndPath := strings.SplitN(u.Path, ":", 2)
	if len(hostAndPath) < 2 {
		//usage()
		fmt.Println("<URL: unix:///path/file.sock:/path>")
		os.Exit(1)
	}

	u.Host = hostAndPath[0]
	u.Path = hostAndPath[1]

	reader := strings.NewReader(data)
	if len(data) > 0 {
		// If there are data the request can't be GET (curl behavior)
		if method == "GET" {
			method = "POST"
		}
		// If data begins with @, it references a file
		if string(data[0]) == "@" && len(data) > 1 {

			buf, err := ioutil.ReadFile(string(data[1:]))
			if err != nil {
				fmt.Println("Failed to open file:", err)
				os.Exit(1)
			}
			reader = strings.NewReader(string(buf))
		}
	}

	query := ""
	if len(u.RawQuery) > 0 {
		query = "?" + u.RawQuery
	}
	req, err := http.NewRequest(method, u.Path+query, reader)
	if err != nil {
		fmt.Println("Fail to create http request", err)
		os.Exit(1)
	}
	if err := addHeaders(req, header); err != nil {
		fmt.Println("Fail to add headers:", err)
		os.Exit(1)
	}
	if err := addCookies(req, cookie); err != nil {
		fmt.Println("Fail to add cookies:", err)
		os.Exit(1)
	}

	conn, err := net.Dial("unix", u.Host)
	if err != nil {
		fmt.Println("Fail to connect to", u.Host, ":", err)
		os.Exit(1)
	}
	client := httputil.NewClientConn(conn, nil)
	//res, err := requestExecute(conn, client, req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Fail to achieve http request over unix socket", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//return string(body)
	return res.StatusCode, string(body)
}

func UnixSocket(method, url, data string) (int, string) {
	return RequestSocket(method, url, data, "Content-Type:application/json", "")
}

func SocketTest() string {
	_, ss := RequestSocket("GET", "unix:///var/run/docker.sock:/containers/json?all=1", "", "Content-Type:application/json", "")
	return "I am Test\n" + ss
}
