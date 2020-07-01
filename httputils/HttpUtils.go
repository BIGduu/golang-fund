package httputils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)


type HttpUtils struct {
	client *http.Client
	request *http.Request
	httpMethod HttpMethod
	header http.Header
	cookie http.Cookie
	url string
	data string
}

type HttpMethod string
const (
	OPTIONS HttpMethod = "OPTIONS"
	GET HttpMethod = "GET"
	POST HttpMethod = "POST"
	PUT HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)


func (this *HttpMethod) ToString() string {
	return string(GET)
}

func (this *HttpUtils) nobodyRequest(url *string) *http.Request {
	split := strings.Split(*url, "?")
	if len(split) == 2 {
		this.url = split[0] + "?" + this.getParseParam(split[1])
	}
	request, _ := http.NewRequest(this.httpMethod.ToString(), this.url, nil)
	return request
}

func (this *HttpUtils) bodyRequest(url, data *string) *http.Request {
	request , _ := http.NewRequest(this.httpMethod.ToString(),*url,strings.NewReader(*data))
	return request
}

func (this *HttpUtils) initRequest() {
	if this.httpMethod == GET {
		this.request = this.nobodyRequest(&this.url)
	} else {
		this.request = this.bodyRequest(&this.url,&this.data)
	}
}

func (this *HttpUtils) AddHeader(name,values string) {
	header := this.request.Header
	header.Add(name,values)
	this.header = header
}

//http请求
func (this *HttpUtils) NewRequest(httpMethod HttpMethod, url, data string) *HttpUtils {
	this.client = &http.Client{}
	this.httpMethod = httpMethod
	this.url = url
	this.data = data
	this.initRequest()
    return this
}

func (this *HttpUtils) DoRequest() *http.Response {
	response, err := this.client.Do(this.request)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}

func (this *HttpUtils) ReadResponseBody(response *http.Response) string {
	all, _ := ioutil.ReadAll(response.Body)
	return string(all)
}

//将get请求的参数进行转义
func (this *HttpUtils)getParseParam(param string) string {
	return url.PathEscape(param)
}
