package main

import (
	"log"
	"net/http"
	"plugin"
	"time"
	 _ "net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		println(">>>>>>>>>>>>>>>>>>>>>>>>>>>Notify")
		<-sc
		println(">>>>>>>>>>>>>>>>>>>>>>>>>>>Stopping")
		os.Exit(1)
	}()
	for {
		p, _ := plugin.Open("./plugin/myplugin.so")
		add, _ := p.Lookup("Add")
		editHandlerOk, _ := p.Lookup("EditHandlerOk")
		addFunc := add.(func(int, int) int)
		println(addFunc(1, 2))
		println(addFunc(2, 2))

		http.HandleFunc("/", editHandlerOk.(func(http.ResponseWriter, *http.Request)))
		if err := http.ListenAndServe(":8181", nil); err != nil {
			log.Fatal(err.Error())
		}

		time.Sleep(1 * time.Second)
		println("Serving")
	}
}
