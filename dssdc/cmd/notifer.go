package cmd

import (
	"encoding/json"
	"log"

	"github.com/ddosakura/ds-watcher-simple-dev/repo"
)

type postParams struct {
	Data string `json:"data"`
}

type NetNotifier struct {
	CallUrl string
}

func newNetNotifier(callUrl string) *NetNotifier {
	return &NetNotifier{
		CallUrl: callUrl,
	}
}

func (n *NetNotifier) Put(notes *repo.Notes) {
	// fmt.Println(notes.ID, notes)
	// ttt := repo.Notes{}
	// fmt.Println(ttt.ID, ttt)

	d, e := json.Marshal(*notes)
	if e != nil {
		mustLog("WARNING", e)
		return
	}
	n.dispatch(&postParams{
		Data: string(d),
	})
}

/*
func (n *NetNotifier) dispatch(params *postParams) {
	b, err := json.Marshal(params)
	if err != nil {
		log.Println("error: json.Marshal n.params. ", err)
		return
	}
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	req, err := http.NewRequest("POST", n.CallUrl, bytes.NewBuffer(b))
	if err != nil {
		log.Println("error: http.NewRequest. ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", "DSSD")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("notifier call failed. err:", err)
		return
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if resp.StatusCode >= 300 {
		// todo retry???
	}
	log.Println("notifier done .")
}
*/
func (n *NetNotifier) dispatch(params *postParams) {
	// b, err := json.Marshal(params)
	// if err != nil {
	// 	log.Println("error: json.Marshal n.params. ", err)
	// 	return
	// }

	// client := &http.Client{
	// 	Timeout: time.Second * 15,
	// }

	// v := url.Values{"data": {params.Data}}.Encode()
	h := NewHttpSend(GetUrlBuild(n.CallUrl, map[string]string{"data": params.Data}))
	_, err := h.Get()
	if err != nil {
		log.Println("notifier call failed. err:", err)
		return
	}

	/*
		client := &http.Client{
			Timeout: time.Second * 15,
		}
		req, err := http.NewRequest("POST", n.CallUrl, strings.NewReader(v))
		if err != nil {
			log.Println("error: http.NewRequest. ", err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
		req.Header.Set("User-Agent", "DSSD")
		resp, err := client.Do(req)
		if err != nil {
			log.Println("notifier call failed. err:", err)
			return
		}
		defer func() {
			if resp != nil && resp.Body != nil {
				_ = resp.Body.Close()
			}
		}()
		if resp.StatusCode >= 300 {
			// todo retry???
		}
		log.Println("notifier done .")
	*/
}
