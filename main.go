package main

import (
	"fmt"
	"plugin"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mooni0929/gowork/intertest"
)

func main() {

	println("go 함수 windows에서")
	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("Hello %d", i)
		go restyRun(100000, msg)
	}
	println("go 함수 window에서 고침 ")

	/*
		varSymbol := getPluginSymbol()
		for i := 1; i <= 100; i++ {
			msg := fmt.Sprintf("Hello %d", i)
			varinterface := getInterfaceHello(varSymbol, msg)
			go runInterface(varinterface, 1000000, msg)
		}
	*/

	time.Sleep(3600 * time.Second)

	/*
			client := resty.New()
			resp, err := client.R().
				EnableTrace().
				Get("http://192.168.0.111:30010")
			if err != nil {
				println(err.Error())
				return
			}


		fmt.Println(resp)
	*/
	println("func End=>")
	time.Sleep(3600 * time.Second)
}
func restyRun(count int, msg string) {
	s := resty.New()
	s.DisableTrace()
	req := s.R()

	println("nestyRun Start=>", msg, "count==>", count)
	for i := 1; i <= count; i++ {
		resp, err := req.
			Get("http://192.168.0.111:30010")

		check := i % 1000
		if check == 0 {

			fmt.Println("Name :", msg, "  count     :", i)

		}

		if err != nil {
			println(err.Error())
			return
		}

		if resp.StatusCode() != 200 {
			fmt.Println("Name :", msg, "  Status     :", resp.StatusCode())
			return
		}
		//time.Sleep(10 * time.Millisecond)

	}
	println("nestyRun end=>", msg)
}
func runInterface(inter intertest.Hello, count int, msg string) {
	for i := 1; i <= count; i++ {
		inter.Run()
		time.Sleep(100 * time.Millisecond)

	}
	println("func end=>", msg)
}
func getPluginSymbol() plugin.Symbol {
	plug, err := plugin.Open("pluginB.so")
	if err != nil {
		println(err.Error())
		return nil
	}
	testSymbol, err := plug.Lookup("GetHttpInterface")
	if err != nil {
		println(err.Error())
		return nil
	}

	return testSymbol

}
func getInterfaceHello(symbol plugin.Symbol, name string) intertest.Hello {
	interHello := symbol.(func(string) intertest.Hello)(name)
	return interHello

}
