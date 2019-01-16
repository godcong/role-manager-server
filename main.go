//go:generate apidoc -i ./service
//go:generate statik -f -src=./doc
package main

import (
	"fmt"
	"github.com/godcong/role-manager-server/service"

	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/godcong/role-manager-server/statik"
)

func main() {
	file, err := os.OpenFile("manager.log", os.O_SYNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Lshortfile)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	service.Start()
	//start
	go func() {
		sig := <-sigs
		//bm.Stop()
		fmt.Println(sig, "exiting")
		done <- true
	}()
	<-done
}
