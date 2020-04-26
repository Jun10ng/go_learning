package github

import (

	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
type errIssue struct {
	err error
	resp *http.Response
}

//获取资源
func (ei * errIssue)get(url string){
	if ei.err!=nil{
		return
	}
	fmt.Printf("search url : %v\n",url)
	ei.resp,ei.err = http.Get(url)
}
//检查状态 && 关闭body
func (ei * errIssue)checkStatue() {
	if  ei.err!=nil{
		return
	}
	if  ei.resp.StatusCode!=http.StatusOK{
		ei.resp.Body.Close()
		ei.err = fmt.Errorf("search query failed %s",ei.resp.Status)
	}
}
// json 解码
func (ei *errIssue)jsonDecode(result interface{}){
	if ei.err !=nil{
		return
	}
	if ei.err = json.NewDecoder(ei.resp.Body).Decode(result);ei.err != nil {
		ei.resp.Body.Close()
	}
}

// 搜索
func SearchIssue(terms []string)(*IssuesSearchResult,error)  {
	q_url :=url.QueryEscape(strings.Join(terms," "))
	ei := errIssue{}
	ei.get(IssuesURL+"?q="+q_url)
	ei.checkStatue()
	var result IssuesSearchResult
	ei.jsonDecode(&result)
	if ei.err !=nil{
		return nil,fmt.Errorf("%w",ei.err)
	}
	ei.resp.Body.Close()
	return &result,nil

}

/*//搜索issue
func SearchIssue(terms []string) (*IssuesSearchResult,error)  {
	q := url.QueryEscape(strings.Join(terms," "))
	resp, err := http.Get(IssuesURL+"?q="+q)
	if err!=nil{
		return nil,err
	}

	//关闭resp.Body
	//使用defer能使得其更简单，第5章会介绍
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result);err!=nil{
		resp.Body.Close()
		return nil,err
	}
	resp.Body.Close()
	return &result,nil
}
*/
