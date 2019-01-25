package uploader

import "encoding/json"

type Dto struct {
	User string
	File string
}

func makeURL(url string, dto Dto) string {
	d, _ := json.Marshal(dto)
	return url + "?data=" + string(d)
}
