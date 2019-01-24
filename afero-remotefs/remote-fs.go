package remoteAfero

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/afero"
)

// Fs is remote filesystem base http(s)
type Fs struct {
	remoteURL  string
	cache      afero.Fs
	lastUpdate map[string]int64
}

// NewRemoteFs create a new RemoteFs
func NewRemoteFs(remoteURL string) afero.Fs {
	return &Fs{
		remoteURL:  remoteURL,
		cache:      afero.NewMemMapFs(),
		lastUpdate: map[string]int64{},
	}
}

// Create creates a file in the filesystem, returning the file and an
// error, if any happens.
func (r *Fs) Create(name string) (afero.File, error) {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Create")
	return nil, nil
}

// Mkdir creates a directory in the filesystem, return an error if any
// happens.
func (r *Fs) Mkdir(name string, perm os.FileMode) error {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Mkdir")
	return nil
}

// MkdirAll creates a directory path and all parents that does not exist
// yet.
func (r *Fs) MkdirAll(path string, perm os.FileMode) error {
	// TODO:
	log.Println("TEMP", "RemoteFs", "MkdirAll")
	return nil
}

func getName(fis []os.FileInfo, err error) ([]string, error) {
	if err != nil {
		return []string{}, err
	}
	res := make([]string, len(fis))
	for k, v := range fis {
		res[k] = v.Name()
	}
	return res, nil
}

// Open opens a file, returning it or an error, if any happens.
func (r *Fs) Open(name string) (afero.File, error) {
	// var (
	// 	root, _ = getName(afero.ReadDir(r.cache, "/"))
	// 	js, _   = getName(afero.ReadDir(r.cache, "/js"))
	// 	css, _  = getName(afero.ReadDir(r.cache, "/css"))
	// 	img, _  = getName(afero.ReadDir(r.cache, "/img"))
	// )
	// fmt.Println(root, js, css, img /*, e1, e2, e3, e4*/)

	// TODO: check file update
	log.Println("TEMP", "RemoteFs", "Open", name)
	f, e := r.cache.Open(name)
	if e == nil {
		fmt.Println(name, "in cache")
		return f, nil
	}
	// e.(*os.PathError).Err == afero.ErrFileNotFound
	if e.(*os.PathError).Err == os.ErrNotExist && !r.nearlyGet(name) {
		r.lastUpdate[name] = time.Now().UnixNano()
		log.Println( /*r.lastUpdate[name], */ "GET", r.remoteURL+"/"+name)
		resp, err := http.Get(r.remoteURL + "/" + name)
		if err != nil {
			fmt.Println(name, "get error", err)
			return nil, err
		}
		if resp.StatusCode == 404 {
			fmt.Println(name, "get 404")
			return nil, err
		}
		file, err := r.cache.Create(name /* + ".tmp"*/)
		if err != nil {
			fmt.Println(name, "cache save error", err)
			return nil, err
		}
		// if resp.ContentLength < 0 {
		// 	fmt.Println(resp)
		// }
		var (
			size     = resp.ContentLength
			dataSize = min(1024, size)
			data     = make([]byte, dataSize)
			n        = dataSize
			// downSize = 0
		)
		defer resp.Body.Close()
		for /*n == dataSize*/ {
			n, err = resp.Body.Read(data)
			file.Write(data[:n])
			// downSize += n
			// fmt.Println(downSize, "/", size)
			if err == io.EOF {
				fmt.Println(name, "get ok")
				file.Close()
				// return file, nil
				return r.Open(name)
			} else if err != nil {
				fmt.Println(name, "cache write error", err)
			}
		}
		// RemoteFile 方案
		// file := NewRemoteFile(r, resp)
		// return file, nil
	}
	fmt.Println(name, "cache open error", e)
	return f, e
}

// OpenFile opens a file using the given flags and the given mode.
func (r *Fs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	// TODO:
	log.Println("TEMP", "RemoteFs", "OpenFile")
	return nil, nil
}

// Remove removes a file identified by name, returning an error, if any
// happens.
func (r *Fs) Remove(name string) error {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Remove")
	return nil
}

// RemoveAll removes a directory path and any children it contains. It
// does not fail if the path does not exist (return nil).
func (r *Fs) RemoveAll(path string) error {
	// TODO:
	log.Println("TEMP", "RemoteFs", "RemoveAll")
	return nil
}

// Rename renames a file.
func (r *Fs) Rename(oldname, newname string) error {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Rename")
	return nil
}

// Stat returns a FileInfo describing the named file, or an error, if any
// happens.
func (r *Fs) Stat(name string) (os.FileInfo, error) {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Stat")
	return nil, nil
}

// The name of this FileSystem
func (r *Fs) Name() string {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Name")
	return ""
}

//Chmod changes the mode of the named file to mode.
func (r *Fs) Chmod(name string, mode os.FileMode) error {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Chmod")
	return nil
}

//Chtimes changes the access and modification times of the named file
func (r *Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	// TODO:
	log.Println("TEMP", "RemoteFs", "Chtimes")
	return nil
}

func (r *Fs) nearlyGet(name string) bool {
	l := r.lastUpdate[name]
	// fmt.Println(l, time.Now().UnixNano())
	if (l == 0) || (time.Now().UnixNano()-l > 1000*1000*1000*1) { // 1s
		return false
	}
	return true
}
