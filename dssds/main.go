package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "../dssdc/statik" // TODO: 考虑放公共部分
	"../repo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rakyll/statik/fs"
	"golang.org/x/net/websocket"
)

func init() {
	repo.Pre = func() *gorm.DB {
		// TODO: 从环境变量里搞进来
		user := os.Getenv("DSSDS_USER")
		if user == "" {
			user = "root"
		}
		pass := os.Getenv("DSSDS_PASS")
		if pass == "" {
			pass = "123456"
		}
		dbHost := os.Getenv("DSSDS_DB_HOST")
		if dbHost == "" {
			dbHost = "127.0.0.1"
		}
		dbPort := os.Getenv("DSSDS_DB_PORT")
		if dbPort == "" {
			dbPort = "3306"
		}
		dbName := os.Getenv("DSSDS_DB_NAME")
		if dbName == "" {
			dbName = "dssd"
		}
		db, err := gorm.Open("mysql", user+":"+pass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			er(err)
		}
		return db
	}
	repo.Init()
}

func main() {
	statikFS, e := fs.New()
	if e != nil {
		fmt.Printf("err: %v\n", e)
		os.Exit(1)
	}
	port := os.Getenv("DSSDS_PORT")
	if port == "" {
		port = ":2000"
	}
	log.Println(port)

	// api
	http.Handle("/developers.action", &apiHandler{api: apiDeveloper})
	http.Handle("/detail.action", &apiHandler{api: apiDetail})
	http.Handle("/note.action", &apiHandler{api: apiNote})

	// websocket fresh
	http.Handle("/fresh", websocket.Handler(wsFreshHandler))

	// view
	http.Handle("/", http.FileServer(statikFS))

	fmt.Printf("Listen %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var (
	WebPages = NewSet()
)

func callFreshWebPage(vv interface{}) {
	v, e := json.Marshal(vv)
	if e != nil {
		log.Println("WARNING", e)
		return
	}
	list := WebPages.List()
	// fmt.Println(list, v)
	for i := range list {
		ws := list[i].(*websocket.Conn)
		if err := websocket.Message.Send(ws, v); err != nil {
			fmt.Println("WARNING", err)
			removeFreshWebPageCallback(ws)
		}
	}
}

func wsFreshHandler(ws *websocket.Conn) {
	defer removeFreshWebPageCallback(ws)
	addFreshWebPageCallback(ws)
	var reply string
	for {
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			break
		}
	}
}

func addFreshWebPageCallback(ws *websocket.Conn) {
	WebPages.Add(ws)
}

func removeFreshWebPageCallback(ws *websocket.Conn) {
	WebPages.Remove(ws)
}
