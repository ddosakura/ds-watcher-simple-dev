package remoteAfero

import (
	"log"
	"net/http"
	"os"

	"github.com/spf13/afero"
)

type File struct {
	remoteFs *Fs
	resp     *http.Response
}

// NewRemoteFile create a remote file obj by remoteFs and response
func NewRemoteFile(remoteFs *Fs, resp *http.Response) afero.File {
	return &File{
		remoteFs,
		resp,
	}
}

func (f *File) Name() string {
	// TODO:
	log.Println("TEMP", "RemoteFile", "Name")
	return ""
}
func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	// TODO:
	log.Println("TEMP", "RemoteFile", "Readdir")
	return nil, nil
}
func (f *File) Readdirnames(n int) ([]string, error) {
	// TODO:
	log.Println("TEMP", "RemoteFile", "Readdirnames")
	return nil, nil
}
func (f *File) Stat() (os.FileInfo, error) {
	log.Println("TEMP", "RemoteFile", "Stat")
	return &FileInfo{f}, nil
}
func (f *File) Sync() error {
	// TODO:
	log.Println("TEMP", "RemoteFile", "Sync")
	return nil
}
func (f *File) Truncate(size int64) error {
	// TODO:
	log.Println("TEMP", "RemoteFile", "Truncate")
	return nil
}
func (f *File) WriteString(s string) (ret int, err error) {
	// TODO:
	log.Println("TEMP", "RemoteFile", "WriteString")
	return -1, nil
}

func (f *File) Close() error {
	// TODO: lock
	log.Println("TEMP", "RemoteFile", "Close")
	return f.resp.Body.Close()
}

func (f *File) Read(b []byte) (n int, err error) {
	// TODO: lock
	log.Println("TEMP", "RemoteFile", "Read")
	return f.resp.Body.Read(b)
}

func (f *File) ReadAt(b []byte, off int64) (n int, err error) {
	// TODO:
	log.Println("TEMP", "RemoteFile", "ReadAt")
	return
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
	// TODO:
	log.Println("TEMP", "RemoteFile", "Seek")
	return -1, nil
}

func (f *File) Write(b []byte) (n int, err error) {
	// TODO:
	log.Println("TEMP", "RemoteFile", "Write")
	return
}

func (f *File) WriteAt(b []byte, off int64) (n int, err error) {
	// TODO:
	log.Println("TEMP", "RemoteFile", "WriteAt")
	return
}
