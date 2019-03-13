//go:generate apidoc -i ./service
//go:generate statik -f -src=./doc
//go:generate protoc --go_out=plugins=grpc:./proto --micro_out=./proto  manager.proto

package main

import (
	"flag"
	"github.com/godcong/go-trait"
	"github.com/godcong/role-manager-server/config"
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/service"
	_ "github.com/godcong/role-manager-server/statik"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "load config file from path")
var elk = flag.Bool("elk", false, "set log to elk")
var logPath = flag.String("log", "logs/manager.log", "set the default log path")

func main() {
	flag.Parse()
	if *elk {
		trait.InitElasticLog("role-manager-server", nil)
	} else {
		trait.InitRotateLog(*logPath, nil)
	}

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
		logrus.Println(sig, "exiting")
		service.Stop()
		done <- true
	}()
	<-done
}
