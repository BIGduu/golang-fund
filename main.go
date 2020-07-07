package main

import (
	"context"
	"fmt"
	"fund/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
	"time"
)

func main() {
	/*	fund := entity.AllFund{}
		go  fund.NewAllFund()
	*/
	responses := make(chan entity.JZResponse)
	jzResponse := entity.JZResponse{}
	go jzResponse.NewJZResponse(responses)
	value, ok := <-responses
	if ok {
		value.Data.List = value.Data.Filter(isAdd)
	}
	list := value.Data.List
	var mongodbURI = "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	if err != nil {
		println(err)
		panic(err)
	}
	_ = client.Connect(ctx)
	now := time.Now()
	toDay := now.Format("2006-01-02 15:04:05")
	collection := client.Database("fund").Collection("isadd" + toDay)
	_, err = collection.InsertMany(ctx, toInterface(list))
	if err != nil {
		fmt.Println("1" + err.Error())
	}
	_, err = collection.InsertMany(ctx, toInterface(list))
	if err != nil {
		fmt.Println("2" + err.Error())
	}
	fmt.Println("执行完毕")
	client.Disconnect(ctx)
}

func toInterface(array interface{}) []interface{} {
	of := reflect.ValueOf(array)
	if of.Kind() == reflect.Slice {
		interfaces := make([]interface{}, of.Len())
		for i := 0; i < of.Len(); i++ {
			interfaces[i] = of.Index(i).Interface()
		}
		return interfaces
	}
	return nil
}

func isAdd(fund entity.JZFund) bool {
	if fund.Gszzl == "---" {
		return false
	}
	if strings.Contains(fund.Gszzl, "-") {
		return false
	}
	return true
}

func isNotAdd(fund entity.JZFund) bool {
	if fund.Gszzl == "---" {
		return false
	}
	if strings.Contains(fund.Gszzl, "-") {
		return true
	}
	return false
}
