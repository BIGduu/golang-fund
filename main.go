package main

import (
	"fund/entity"
	"gopkg.in/mgo.v2"
	"strings"
)

func main() {
/*	fund := entity.AllFund{}
	go  fund.NewAllFund()
*/
	session, err := mgo.Dial("mongodb://127.0.0.1:27027")
	if err != nil {
		panic(err)
	}
	responses := make(chan entity.JZResponse)
	jzResponse := entity.JZResponse{}
	go  jzResponse.NewJZResponse(responses)
	value ,ok:= <- responses
	if ok {
		value.Data.Filter(isAdd)
	}
	defer session.Close()
	c := session.DB("test").C("fund")
	_ = c.Insert(value.Data.List)

}

func isAdd(fund entity.JZFund) bool{
	if strings.Contains(fund.Gszzl,"-") {
		return false
	}
	return true
}
