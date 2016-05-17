package main

import (
	_ "bapi/routers"
	"github.com/astaxie/beego"
	_ "bapi/models"
	"gopkg.in/mgo.v2"
	"os"
	"log"
)

func main() {
	mgo.SetDebug(true)
	logout := log.New(os.Stdout, "MGO: ", log.Lshortfile)
	mgo.SetLogger(logout)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
