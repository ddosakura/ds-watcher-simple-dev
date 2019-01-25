package main

// TODO: 这也是从dssdc复制来的，相同部分考虑做放公共部分

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ddosakura/ds-watcher-simple-dev/repo"
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
	r.ParseForm()
	data, e := h.api(r)
	w.Header().Set("Content-Type", "application/json;chartset=uft-8")
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

func apiDeveloper(*http.Request) (interface{}, error) {
	return repo.Developers(), nil
}

func apiDetail(r *http.Request) (interface{}, error) {
	name := r.FormValue("name")
	// fmt.Println(r.Form, name)
	return *repo.Detail(name), nil
}

func apiNote(r *http.Request) (interface{}, error) {
	// result, _ := ioutil.ReadAll(r.Body)
	// r.Body.Close()
	// fmt.Printf("%s\n", result)
	// var f interface{}
	// json.Unmarshal(result, &f)
	// fmt.Println(f)
	// m := f.(map[string]interface{})

	// fmt.Println(r, r.Method)
	// fmt.Println(r.Form)
	data := r.FormValue("data")
	d := repo.Notes{}
	json.Unmarshal([]byte(data), &d)
	// fmt.Println(d)
	repo.Note(&d)
	/*
		repo.Note(&repo.Notes{
			Developer: r.FormValue(""),
			Project:   r.FormValue(""),
			File:      r.FormValue(""),
			Ext:       r.FormValue(""),
			// ChangeTime: nil,
		})
	*/
	return "", nil
}
