package main

import (
	server "github.com/itimofeev/vhustle/modules"
	"github.com/itimofeev/vhustle/modules/util"

	"log"
	"net/http"
)

func main() {
	util.InitPersistence()
	server.InitCronTasks()
	util.AnyLog.Debug("Lets start fun with vhustle :)")

	log.Fatal(http.ListenAndServe(":8080", server.InitRouter()))
}
