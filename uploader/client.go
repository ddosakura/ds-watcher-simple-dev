package uploader

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Upload(filename string, url string, dto Dto) {
	file, err := os.Open(filename)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	// s, e := file.Stat()
	// fmt.Println(e, s.Size())

	res, err := http.Post(makeURL(url, dto), "binary/octet-stream", file)
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}
