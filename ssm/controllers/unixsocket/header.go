package unixsocket

import (
	"fmt"
	"net/http"
	"strings"
)

type Headers map[string]string

func addHeaders(req *http.Request, headerstr string) error {
	headers, err := parseHeader(headerstr)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return nil
}

func parseHeader(headerstr string) (Headers, error) {
	headers := Headers{}
	if len(headerstr) == 0 {
		return nil, nil
	}
	headersStr := strings.Split(headerstr, "|")
	for _, headerStr := range headersStr {
		headerArr := strings.Split(headerStr, ":")
		if len(headerArr) != 2 {
			return nil, fmt.Errorf("Invalid header:", headerStr)
		}
		headers[strings.TrimSpace(headerArr[0])] = strings.TrimSpace(headerArr[1])
	}
	return headers, nil
}
