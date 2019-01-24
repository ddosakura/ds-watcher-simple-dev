package remoteAfero

import (
	"log"
	"os"
	"time"
)

type FileInfo struct {
	f *File
}

// base name of the file
func (fi *FileInfo) Name() string {
	// TODO:
	log.Println("TEMP", "RemoteFileInfo", "Name")
	return ""
}

// length in bytes for regular files; system-dependent for others
func (fi *FileInfo) Size() int64 {
	// TODO:
	log.Println("TEMP", "RemoteFileInfo", "Size", fi.f.resp.ContentLength)
	return fi.f.resp.ContentLength
}

// file mode bits
func (fi *FileInfo) Mode() os.FileMode {
	// TODO:
	log.Println("TEMP", "RemoteFileInfo", "Mode")
	return 0
}

// modification time
func (fi *FileInfo) ModTime() time.Time {
	// TODO:
	log.Println("TEMP", "RemoteFileInfo", "ModTime")
	return time.Now()
}

// abbreviation for Mode().IsDir()
func (fi *FileInfo) IsDir() bool {
	// TODO:
	log.Println("TEMP", "RemoteFileInfo", "IsDir")
	return false
}

// underlying data source (can return nil)
func (fi *FileInfo) Sys() interface{} {
	// TODO:
	log.Println("TEMP", "RemoteFileInfo", "Sys")
	return false
}
