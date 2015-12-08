package controllers

import (
	//"fmt"
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

var (
	MasterAddr string
	MasterPort string
	passAuth   string
	innerPort  string
	serverAddr string
)

func init() {
	data := ReadList("conf/app.conf")
	MasterAddr = data["masteraddr"]
	MasterPort = data["masterport"]
	passAuth = data["passauth"]
	innerPort = data["innerport"]
	serverAddr = data["server"]
}

//List all the configuration file
func ReadList(filepath string) map[string]string {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var data map[string]string
	data = make(map[string]string)
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			if len(section) > 0 {
				data[section+"::"+strings.TrimSpace(line[0:i])] = value
			} else {
				data[strings.TrimSpace(line[0:i])] = value
			}

		}

	}

	return data
}
