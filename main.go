//go:generate apidoc -i ./service
//go:generate statik -f -src=./doc
//go:generate protoc --go_out=plugins=grpc:./proto manager.proto
package main

import (
	"flag"
	"fmt"
	"github.com/godcong/go-trait"
	"github.com/godcong/role-manager-server/config"
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/service"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/godcong/role-manager-server/statik"
)

var configPath = flag.String("path", "config.toml", "load config file from path")

func main() {
	flag.Parse()
	trait.InitElasticLog("role-manager-server", nil)
	err := config.Initialize(*configPath)
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
