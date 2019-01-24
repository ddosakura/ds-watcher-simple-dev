package cmd

import (
	"compress/flate"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

func pkg(tgName string) {
	os.Remove(tgName)

	iDirs := cfg.Monitors.GetStringSlice("includeDirs")
	// TODO: 过滤忽略的文件
	// eDirs := cfg.Monitors.GetStringSlice("exceptDirs")
	// allFile := make([]string, 0)
	// for _, v := range iDirs {
	// 	allFile = walkArchive(v, &eDirs, allFile)
	// }
	tg := archiver.NewTarGz()
	tg.CompressionLevel = flate.BestCompression
	// tg.Tar.MkdirAll = true
	// e := tg.Archive(allFile, tgName)
	e := tg.Archive(iDirs, tgName)
	if e != nil {
		er(e)
	}
}

func walkArchive(path string, eDirs *[]string, all []string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		er(err)
	}
	dirs := make([]string, 0)
	for _, v := range files {
		p := filepath.Join(path, v.Name())
		if v.IsDir() {
			if !inStringArray(v.Name(), eDirs) {
				dirs = append(dirs, p)
			}
		} else {
			mustLog("archive", p)
			// e := tg.Archive([]string{p}, tgName)
			// if e != nil {
			// 	er(e)
			// }
			all = append(all, p)
		}
	}
	for _, v := range dirs {
		all = walkArchive(v, eDirs, all)
	}
	return all
}
