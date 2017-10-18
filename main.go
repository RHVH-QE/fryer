package main

import (
	"github.com/dracher/fryer/autocore"
	"github.com/dracher/fryer/model"
	"github.com/dracher/fryer/server"
	"github.com/dracher/fryer/utils/cobbler"
)

var q = model.NewQuery()

func main() {
	hosts, _ := q.Hosts()
	autocore.InitScheduler(hosts)
	cobbler.InitCobblerConfig(q)
	server.InitApp(true, q).Run(":8090")
	// job := autocore.NewJob("manual")
	// log.Info(job)
	// job.Event("preJobCobbler", "dell501-01", "RHVH-73-4.1-20170120.4", "managed by fryer", "testing", "gggg", []string{"macaddress-enp2s0", "00:22:19:27:54:c7"})
	// log.Info(job.Current())
	// log.Info(job.Can("preJobBeaker"))
}
