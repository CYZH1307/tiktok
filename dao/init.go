package dao

import (
	"encoding/json"
	"github.com/CYZH1307/tiktik/config"
	"gorm.io/driver/mysql"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var DB *gorm.DB

func init() {
	cfg := comfig.DB
	conn := cfg["user"] + ":" + cfg["password"] + "@tcp(" + cfg["host"] + ":" + cfg["port"] + ")/" + cfg["dbname"] + "?charset=utf8mb4&parseTime=true"
	DB, err = gorm.Open(mysql.Open(conn))
	Handle(err)

	if config.Video["video_prefix"] == "" {
		config.Video["video_prefix"] = "http://" + getIP() + config.Port + "/douyin/video/"
	}
	if config.Video["cover_prefix"] == "" {
		config.Video["cover_prefix"] = "http://" + getIP() + config.Port + "/douyin/cover/"
	}
}

func Handle(e error) {
	if e != nil {
		log.Panicf("[ERR] Tiktok DAO Layer Error: %v", e)
	}
}

func getIP() string {
	if config.NetEnv == "internal" {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		Handle(err)
		defer func(conn net.Conn) {
			err = conn.Close()
			Handle(err)
		}(conn)
		localAddr := conn.LocalAddr().(*net.UDPAddr).String()
		addr, _, err := net.SplitHostPort(localAddr)
		return addr
	}

	type IP struct {
		Query string
	}
	req, err := http.Get("http://ip-api.com/json/")
	Handle(err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		Handle(err)
	}(req.Body)
	body, err := ioutil.ReadAll(req.Body)
	Handle(err)
	var ip IP
	err = json.Unmarshal(body, &ip)
	Handle(err)
	return ip.Query
}

