package cmd

import (
	"log"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	watcher *fsnotify.Watcher

	taskMan *TaskMan

	exts  []string
	delay int
	cmds  []string
)

func eventDispatcher(event fsnotify.Event) {
	ext := path.Ext(event.Name)
	if len(exts) > 0 &&
		exts[0] != ".*" &&
		!inStringArray(ext, &exts) {
		log.Println(ext, exts, inStringArray(ext, &exts))
		return
	}
	switch event.Op {
	case
		fsnotify.Write,
		fsnotify.Rename:
		log.Println("EVENT", event.Op.String(), event.Name)
		taskMan.Put(&changedFile{
			Name:    event.Name,
			Changed: time.Now().UnixNano(),
			Ext:     ext,
		})
	case fsnotify.Remove:
	case fsnotify.Create:
	}
}

func addWatcher() {
	log.Println("loading directory...")
	iDirs := cfg.Monitors.GetStringSlice("includeDirs")
	eDirs := cfg.Monitors.GetStringSlice("exceptDirs")
	// log.Println(iDirs)
	dirs := make([]string, 0)
	for _, v := range iDirs {
		dirs = append(dirs, v)
		dirs = appendDirWatcher(dirs, v, &eDirs)
	}
	for _, dir := range dirs {
		log.Println("watcher add -> ", dir)
		err := watcher.Add(dir)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("ready!")
}

func initWatcher() {
	exts = cfg.Monitors.GetStringSlice("types")
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	delay = cfg.Commands.GetInt("delaymillsecond")
	cmds = cfg.Commands.GetStringSlice("exec")
	done := make(chan bool)
	taskMan = newTaskMan(cfg.Commands.GetInt("DelayMillSecond"))
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				eventDispatcher(event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	addWatcher()
	<-done
}
