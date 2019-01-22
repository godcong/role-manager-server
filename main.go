//go:generate apidoc -i ./service
//go:generate statik -f -src=./doc
//go:generate protoc --go_out=plugins=grpc:./proto manager.proto
package main

import (
	"flag"
	"fmt"
	"github.com/godcong/role-manager-server/config"
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/service"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/godcong/role-manager-server/statik"
)

var configPath = flag.String("path", "config.toml", "load config file from path")

func main() {
	flag.Parse()
	file, err := os.OpenFile("manager.log", os.O_SYNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Lshortfile)

	err = config.Initialize(*configPath)
	if err != nil {
		panic(err)
	}

	model.InitDB(config.Config())

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//start
	service.Start()

	go func() {
		sig := <-sigs
		fmt.Println(sig, "exiting")
		service.Stop()
		done <- true
	}()
	<-done
}
