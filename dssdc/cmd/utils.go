package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func er(msg interface{}) {
	log.Panicln(msg)
}

var wd string

func initWD() {
	tempWD, err := os.Getwd()
	if err != nil {
		er(err)
	}
	wd = tempWD
}

func getPath(arg string) string {
	if cfgFile[0] == '.' {
		initWD()
		arg = filepath.Join(wd, arg)
	}
	return arg
}

func getRootPath() string {
	return getPath(cfgFile)
}

func getConfigPath() string {
	return filepath.Join(getRootPath(), "dssdc.yaml")
}

func getDBPath(db string) string {
	arg := db
	if cfgFile[0] == '.' {
		arg = filepath.Join(getRootPath(), arg)
	}
	return arg
}

func appendDirWatcher(dirs []string, path string, eDirs *[]string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		er(err)
	}
	for _, v := range files {
		if v.IsDir() {
			p := filepath.Join(path, v.Name())
			if !inStringArray(v.Name(), eDirs) {
				dirs = append(dirs, p)
				dirs = appendDirWatcher(dirs, p, eDirs)
			}
		}
	}
	return dirs
}

func inStringArray(value string, arr *[]string) bool {
	for _, v := range *arr {
		if value == v {
			return true
		}
	}
	return false
}

func cmdParse2Array(s string, cf *changedFile) []string {
	a := strings.Split(s, " ")
	r := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if ss := strings.Trim(a[i], " "); ss != "" {
			r = append(r, strParseRealStr(ss, cf))
		}
	}
	return r
}

func strParseRealStr(s string, cf *changedFile) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(s, "{{file}}", cf.Name, -1),
			"{{ext}}", cf.Ext, -1,
		),
		"{{changed}}", strconv.FormatInt(cf.Changed, 10), -1,
	)
}

func mustLog(v ...interface{}) {
	log.Println(v...)
	if !pLog {
		fmt.Println(v...)
	}
}
