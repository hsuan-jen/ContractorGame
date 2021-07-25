package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"
)

var Server struct {
	LogLevel        string
	LogPath         string
	WSAddr          string
	CertFile        string
	KeyFile         string
	TCPAddr         string
	MaxConnNum      int
	ConsolePort     int
	ProfilePath     string
	ConfPath        string
	MgodbAddr       string
	GameMgoConnNum  int
	LoginMgoConnNum int
}

func init() {
	data, err := ioutil.ReadFile("../../bin/conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
