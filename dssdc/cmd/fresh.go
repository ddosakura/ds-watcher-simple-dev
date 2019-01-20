package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	_ "../statik"
	"github.com/rakyll/statik/fs"
	"golang.org/x/net/websocket"
)

var openCommands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func initFreshing() {
	statikFS, e := fs.New()
	if e != nil {
		fmt.Printf("err: %v\n", e)
		os.Exit(1)
	}
	port := os.Getenv("WSDB_PORT")
	if port == "" {
		port = ":" + strconv.Itoa(cfg.Port)
	}
	log.Println(port)

	// websocket fresh
	http.Handle("/fresh", websocket.Handler(wsFreshHandler))
	// websocket
	http.Handle("/ws", websocket.Handler(wsHandler))

	// view
	http.Handle("/", http.FileServer(statikFS))

	// TODO: 优化（考虑第三方库）
	// open the web browser
	run, ok := openCommands[runtime.GOOS]
	if !ok {
		mustLog("WARNING", "don't know how to open things on %s platform", runtime.GOOS)
	}
	cmd := exec.Command(run, "http://localhost"+port)
	go cmd.Start()

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

// TODO: remove test ws
func wsHandler(ws *websocket.Conn) {
	defer removeConn(ws)
	addConn(ws)
	for {
		var reply string
		//websocket接受信息
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}
		go broadcastConn(reply)
	}
}

var (
	connGroup = NewSet()
)

func addConn(ws *websocket.Conn) {
	connGroup.Add(ws)
}

func removeConn(ws *websocket.Conn) {
	connGroup.Remove(ws)
}

func broadcastConn(msg string) {
	list := connGroup.List()
	for i := range list {
		if err := websocket.Message.Send(list[i].(*websocket.Conn), msg); err != nil {
			fmt.Println("send failed:", err)
		}
	}
}
