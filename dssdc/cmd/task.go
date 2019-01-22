package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/ddosakura/ds-watcher-simple-dev/repo"
)

type changedFile struct {
	Name    string
	Changed int64
	Ext     string
}

type TaskMan struct {
	lastTaskId int64
	delay      int
	cmd        *exec.Cmd
	notifier   *NetNotifier
	putLock    sync.Mutex
	runLock    sync.Mutex
}

func newTaskMan(delay int) *TaskMan {
	tm := &TaskMan{
		delay: delay,
	}
	root := cfg.APIs.GetString("root")
	if root != "" {
		callUrl := cfg.APIs.GetString("notifier")
		if callUrl == "" {
			callUrl = root + "/note.action"
		} else if !(strings.HasPrefix(callUrl, "http://") || strings.HasPrefix(callUrl, "https://")) {
			callUrl = root + "/" + callUrl
		}
		tm.notifier = newNetNotifier(callUrl)

		// TODO: add upload api
	}
	return tm
}

func (t *TaskMan) Put(cf *changedFile) {
	if t.delay < 1 {
		t.preRun(cf)
		return
	}
	t.putLock.Lock()
	defer t.putLock.Unlock()
	t.lastTaskId = cf.Changed
	go func() {
		<-time.Tick(time.Millisecond * time.Duration(t.delay))
		if t.lastTaskId > cf.Changed {
			return
		}
		t.preRun(cf)
	}()
}

func (t *TaskMan) preRun(cf *changedFile) {
	if t.cmd != nil && t.cmd.Process != nil {
		err := t.cmd.Process.Kill()
		if err != nil {
			log.Println("err: ", err)
		}
		log.Println("stop old process ")
	}
	go t.run(cf)
}

func (t *TaskMan) run(cf *changedFile) {
	notes := &repo.Notes{
		Developer:  cfg.Developer,
		Project:    cfg.ProjectName,
		File:       cf.Name,
		Ext:        cf.Ext,
		ChangeTime: time.Now(),
	}

	// TODO: temp data
	// call localdb
	// fmt.Println(cfg.LocalDB)
	if strings.HasSuffix(cfg.LocalDB, ".db") {
		go repo.Note(notes)
	}

	log.Println("fresh webpage")
	go callFreshWebPage(notes)

	// TODO: call web api
	if t.notifier != nil {
		go t.notifier.Put(notes)
	}

	// call cmd
	t.runLock.Lock()
	defer t.runLock.Unlock()
	for i := 0; i < len(cmds); i++ {
		carr := cmdParse2Array(cmds[i], cf)
		log.Println("EXEC", carr)
		t.cmd = exec.Command(carr[0], carr[1:]...)
		//cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_UNICODE_ENVIRONMENT}
		t.cmd.Stdin = os.Stdin
		//cmd.Stdout = os.Stdout
		t.cmd.Stderr = os.Stderr
		t.cmd.Dir = wd
		t.cmd.Env = os.Environ()
		stdout, err := t.cmd.StdoutPipe()
		if err != nil {
			log.Println("error=>", err.Error())
			return
		}
		err = t.cmd.Start()
		if err != nil {
			log.Println("run command", carr, "error. ", err)
		}
		reader := bufio.NewReader(stdout)
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Print(line)
		}
		err = t.cmd.Wait()
		if err != nil {
			log.Println("cmd wait err ", err)
			break
		}
		if t.cmd.Process != nil {
			if err = t.cmd.Process.Kill(); err != nil {
				log.Println("cmd cannot kill ", err)
			}
		}
	}
	log.Println("end ")
}
