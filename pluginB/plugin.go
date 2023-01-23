package main

import (
	"fmt"
	"gowork/intertest"

	"github.com/go-resty/resty/v2"
)

type PlugStruct struct {
	Name       string
	RestClient *resty.Client
}

func (s *PlugStruct) Init() {
	s.RestClient = resty.New()
}
func (s *PlugStruct) Run() {

	if s.RestClient == nil {
		s.Init()
		fmt.Println("init client", s.Name, "client", &s.RestClient)
	}
	resp, err := s.RestClient.R().
		EnableTrace().
		Get("http://192.168.0.111:30010")

	if err != nil {
		println(err.Error())
		return
	}
	if resp.StatusCode() != 200 {
		fmt.Println("Name :", s.Name, "  Status     :", resp.StatusCode())
		return
	}

}
func GetHttpInterface(name string) intertest.Hello {
	s := PlugStruct{Name: name}

	h := intertest.Hello(&s)
	return h
}
