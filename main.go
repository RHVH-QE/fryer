package main

import (
	"github.com/dracher/fryer/autocore"
	"github.com/dracher/fryer/model"
	"github.com/dracher/fryer/server"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

var q = model.NewQuery()

func main() {
	hosts, _ := q.Hosts()
	autocore.InitScheduler(hosts)
	server.InitApp(true, q).Run(":8090")
}
