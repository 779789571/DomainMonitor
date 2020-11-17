package notification

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type serverJiangResult struct {
	Errno   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
	Dataset string `json:"dataset"`
}

func PostToServerJiang(title string, content string, api string) bool {
	//work find
	//title = url.QueryEscape(title)
	//content = url.QueryEscape(content)
	//serverJiangUrl := "https://sc.ftqq.com/" + api + ".send"
	//serverJiangUrl = serverJiangUrl + "?text="+title+"&desp="+content
	//
	//resp, err2 := http.Get(serverJiangUrl)
	//if err2 != nil {
	//	return false
	//}
	//end work fine
	//post
	serverJiangUrl := "https://sc.ftqq.com/" + api + ".send"
	//title = url.QueryEscape(title) #post发送无需urlencode
	//content = url.QueryEscape(content)#post发送无需urlencode
	resp , err2 := http.PostForm(serverJiangUrl, url.Values{"text": {title}, "desp": {content}})
	if err2 != nil {
		return false
	}
	//end post
	defer resp.Body.Close()
	var res serverJiangResult
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)
	//Printf("serverjiang_resp:%s,%s",strconv.Itoa(res.Errno),res.Errmsg)
	if res.Errno == 0 {
		return true
	} else {
		return false
	}
}
