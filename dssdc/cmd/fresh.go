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

	"github.com/ddosakura/ds-watcher-simple-dev/afero-remotefs"
	_ "github.com/ddosakura/ds-watcher-simple-dev/dssdc/statik"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/afero"
	"golang.org/x/net/websocket"
)

var openCommands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var (
	routeMap map[string]string
)

func initFreshing() {
	statikFS, e := fs.New()
	if e != nil {
		fmt.Printf("err: %v\n", e)
		os.Exit(1)
	}
	port := os.Getenv("DSSDC_PORT")
	if port == "" {
		port = ":" + strconv.Itoa(cfg.Port)
	}
	log.Println(port)

	if cfg.Monitors.GetBool("useWebPage") {
		// static
		iDirs := cfg.Monitors.GetStringSlice("includeDirs")
		routeMap = make(map[string]string, len(iDirs))
		for _, v := range iDirs {
			p := 0
			for v[p] == '.' || v[p] == '/' {
				p++
			}
			t := getPath(v)
			// 不加`/`无法获取index.html以外的文件
			v = v[p:] + "/"
			routeMap["/app/"+v] = t

			// fmt.Println("/app/" + v)

			bp := afero.NewBasePathFs(afero.NewOsFs(), t)
			// fmt.Println(afero.ReadDir(bp, ""))
			// fmt.Println(afero.ReadDir(bp, "l"))

			// 标准方案
			httpFs := afero.NewHttpFs(bp)
			// fileserver := http.FileServer(httpFs.Dir("/"))
			fileserver := http.FileServer(httpFs)
			http.Handle("/app/"+v, http.StripPrefix("/app/"+v, fileserver))

			// 新方案，基于statik库改的
			// binding, e := aFsBinding.New(bp)
			// if e != nil {
			// 	mustLog("WARNING", e)
			// 	continue
			// }
			// http.Handle("/app/"+v, http.StripPrefix("/app/"+v, http.FileServer(binding)))

			// 老方案 基于http包改的
			// http.Handle("/app/"+v, http.StripPrefix("/app/"+v, aFsBinding.FileServer(bp)))

			// 原生
			// http.Handle("/app/"+v, http.StripPrefix("/app/"+v, http.FileServer(http.Dir(v))))
		}
	}

	// proxy url
	for localURL, remoteURL := range cfg.Proxy {
		fmt.Println(localURL, remoteURL)
		bp := remoteAfero.NewRemoteFs(remoteURL)
		httpFs := afero.NewHttpFs(bp)
		fileserver := http.FileServer(httpFs)
		http.Handle(localURL, http.StripPrefix(localURL, fileserver))
	}

	// api
	http.Handle("/entrypoint.action", &apiHandler{api: apiEntryPoint})
	http.Handle("/developers.action", &apiHandler{api: apiDeveloper})
	http.Handle("/detail.action", &apiHandler{api: apiDetail})

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
