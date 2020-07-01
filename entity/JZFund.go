package entity

import (
	"encoding/json"
	"fmt"
	"fund/httputils"
	"strings"
)

const JZFundURL = "http://api.fund.eastmoney.com/FundGuZhi/GetFundGZList?type=1&sort=3&orderType=desc&canbuy=1&pageIndex=1&pageSize=20000&callback=jQuery183045152404206471175_1593332615849&_=1593332692698"

type JZFund struct {
	Discount    float32     `json:"Discount"`
	FScaleType  string      `json:"FScaleType"`
	FType       string      `json:"FType"`
	IsExchg     string      `json:"IsExchg"`
	IsListTrade string      `json:"IsListTrade"`
	Jjgsid      string      `json:"JJGSID"` //基金编号
	ListTexch   string      `json:"ListTexch"`
	PLevel      float32     `json:"PLevel"`
	Rate        string      `json:"Rate"`
	Bzdm        string      `json:"bzdm"`
	Dwjz        string      `json:"dwjz"`
	Feature     string      `json:"feature"`
	Fundtype    string      `json:"fundtype"`
	Gbdwjz      string      `json:"gbdwjz"`
	Gspc        string      `json:"gspc"`
	Gsz         string      `json:"gsz"`   //估算值
	Gszzl       string      `json:"gszzl"` //估算净值
	Gszzlcolor  string      `json:"gszzlcolor"`
	Gxrq        string      `json:"gxrq"`  //估算日期
	Gzrq        string      `json:"gzrq"`  //估值日期
	Isbuy       string      `json:"isbuy"` //是否能买
	Jjjc        string      `json:"jjjc"`  //基金名字
	Jjjcpy      string      `json:"jjjcpy"`
	Jjlx        interface{} `json:"jjlx"`
	Jjlx2       interface{} `json:"jjlx2"`
	Jjlx3       interface{} `json:"jjlx3"`
	Jzzzl       string      `json:"jzzzl"`
	Jzzzlcolor  string      `json:"jzzzlcolor"`
	Sgzt        string      `json:"sgzt"`
	Shzt        interface{} `json:"shzt"`
}



type JZBody struct {
	Canbuy   string   `json:"canbuy"`
	Gxrq     string   `json:"gxrq"`
	Gzrq     string   `json:"gzrq"`
	List     []JZFund `json:"list"`
	Sort     string   `json:"sort"`
	SortType string   `json:"sortType"`
	TypeStr  string   `json:"typeStr"`
}

func (this *JZBody) Filter(fun func(fund JZFund) bool){
	list := this.List
	for index, jzFund := range list {
		if fun(jzFund) {
			list = append(list[:index], list[index+1:]...)
		}
	}
	this.List = list
}

type JZResponse struct {
	Data       JZBody      `json:"Data"`
	ErrCode    int64       `json:"ErrCode"`
	ErrMsg     interface{} `json:"ErrMsg"`
	Expansion  interface{} `json:"Expansion"`
	PageIndex  int64       `json:"PageIndex"`
	PageSize   int64       `json:"PageSize"`
	TotalCount int64       `json:"TotalCount"`
	jsonString string
}

func (this *JZResponse) Json() string {
	return this.jsonString
}

func (this *JZResponse) ToJson(responseBody string) {
	start := strings.Index(responseBody, "(") + 1
	end := len(responseBody) - 1
	this.jsonString = responseBody[start:end]
}

func (this *JZResponse) NewJZResponse(sync chan JZResponse) {
	utils := httputils.HttpUtils{}
	request := utils.NewRequest(httputils.GET, JZFundURL, "")
	request.AddHeader("Referer","http://fund.eastmoney.com/fundguzhi.html")
	response := request.DoRequest()
	body := request.ReadResponseBody(response)
	this.ToJson(body)
	err := json.Unmarshal([]byte(this.jsonString), this)
	if err != nil {
		fmt.Println(err)
	}
	sync <- *this
}
