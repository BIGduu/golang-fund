package entity

import (
	"encoding/json"
	"fmt"
	"fund/httputils"
	"reflect"
	"strings"
	"time"
)

const allFunds = "http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=all&rs=&gs=0&sc=zzf&st=desc&sd=%s&ed=%s&qdii=&tabSubtype=,,,,,&pi=1&pn=10000&dx=1&v=0.9522313665975868"

type AllFund struct {
	Datas      []string `json:"datas"`
	allRecords int32    `json:"allRecords,omitempty"`
	pageIndex  int32    `json:"pageIndex,omitempty"`
	pageNum    int32    `json:"pageNum,omitempty"`
	allNum     int32    `json:"allNum,omitempty"`
	allPages   int32    `json:"allPages,omitempty"`
	gpNum      int32    `json:"gpNum,omitempty"`
	hhNum      int32    `json:"hhNum,omitempty"`
	zqNum      int32    `json:"zqNum,omitempty"`
	zsNum      int32    `json:"zsNum,omitempty"`
	bbNum      int32    `json:"bbNum,omitempty"`
	qdiiNum    int32    `json:"qdiiNum,omitempty"`
	etfNum     int32    `json:"etfNum,omitempty"`
	lofNum     int32    `json:"lofNum,omitempty"`
	fofNum     int32    `json:"fofNum,omitempty"`
	startDate  string
	endDate    string
}


func (this *AllFund) getURL() string {
	now := time.Now()
	toDay := now.Format("2006-01-02")
	return fmt.Sprintf(allFunds,toDay,toDay)
}

func(this *AllFund) NewAllFund() {
	utils := httputils.HttpUtils{}
	request := utils.NewRequest(httputils.GET, this.getURL(), "")
	request.AddHeader("Referer","http://fund.eastmoney.com/fundguzhi.html")
	doRequest := request.DoRequest()
	body := request.ReadResponseBody(doRequest)
	index := strings.Index(body, "datas")
	lastIndex := strings.LastIndex(body, ";")
	jsonString := body[index:lastIndex]
	jsonString = "{" + jsonString
	toJson := this.toJson(jsonString)
	err := json.Unmarshal([]byte(toJson), this)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *AllFund) toJson(originJson string) string {
	of := reflect.TypeOf(*this)
	numField := of.NumField()
	for i := 0; i < numField; i++ {
		field := of.Field(i)
		name := field.Name
		name = strings.ToLower(name[0:1]) + name[1:]
		newNamed := "\"" + name + "\""
		originJson = strings.Replace(originJson, name, newNamed, 1)
	}
	return originJson
}
