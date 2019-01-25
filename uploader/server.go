package uploader

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver"
)

func Handler(r *http.Request) (interface{}, error) {
	// fmt.Println(r)
	if r.Method != "POST" {
		return nil, errors.New("Need POST")
	}
	dto := Dto{}
	json.Unmarshal([]byte(r.FormValue("data")), &dto)

	tmpName := "./" + dto.User + "-" + dto.File + "-" + GetRandomSalt() + ".tar.gz"

	file, err := os.Create(tmpName)
	if err != nil {
		return nil, err
	}
	// n, err := io.Copy(file, r.Body)
	_, err = io.Copy(file, r.Body)
	// fmt.Println("r", n)
	if err != nil {
		return nil, err
	}
	file.Close()

	ws := os.Getenv("DSSDS_WORKSPACE")
	target := filepath.Join(ws, dto.User, strings.Split(dto.File, ".")[0])
	// unarchiv
	tg := archiver.NewTarGz()
	tg.OverwriteExisting = true
	err = tg.Unarchive(tmpName, target)
	if err != nil {
		return nil, err
	}
	return dto.User + ": " + dto.File + " upload success", nil
}
