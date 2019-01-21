package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../../repo"
)

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var successMSg = "SUCCESS"

type apiHandler struct {
	api func(r *http.Request) (interface{}, error)
}

func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;chartset=uft-8")
	r.ParseForm()
	data, e := h.api(r)
	var rep CommonResponse
	if e == nil {
		rep = CommonResponse{
			Code: 0,
			Msg:  successMSg,
			Data: data,
		}
	} else {
		rep = CommonResponse{
			Code: -1,
			Msg:  e.Error(),
			Data: data,
		}
	}
	d, e := json.Marshal(rep)
	// fmt.Println(rep)
	if e != nil {
		d, _ := json.Marshal(CommonResponse{
			Code: -1,
			Msg:  e.Error(),
			Data: nil,
		})
		fmt.Fprintln(w, string(d))
		return
	}
	fmt.Fprintln(w, string(d))
}

func apiEntryPoint(*http.Request) (interface{}, error) {
	return routeMap, nil
}

func apiDeveloper(*http.Request) (interface{}, error) {
	return repo.Developers(), nil
}

func apiDetail(r *http.Request) (interface{}, error) {
	name := r.FormValue("name")
	// fmt.Println(r.Form, name)
	return *repo.Detail(name), nil
}