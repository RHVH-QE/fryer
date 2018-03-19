package main

import (
	"flag"
	"fmt"

	"github.com/dracher/fryer/utils/cobbler"
)

var cb *cobbler.Cobbler

func init() {
	cb = cobbler.NewCobblerManual()
}

func newSystem(sysName, profileName, comment, status, kargs, nn, nm string) {
	nics := []string{fmt.Sprintf("macaddress-%s", nn), nm}
	cb.NewSystem(sysName, profileName, comment, status, kargs, nics)
}

func removeSystem(sysName string) {
	cb.RemoveSystem(sysName)
}

func rebootSystem(sysName string) {
	cb.RemoveSystem(sysName)
}

func main() {
	profile := flag.String("p", "RHVH-4.1-74-20180307.5", "the pxe profile name")
	machine := flag.String("m", "hp-z220-16.qe.lab.eng.nay.redhat.com", "the full beaker name")
	nicName := flag.String("nn", "eno1", "nic name")
	nicMac := flag.String("nm", "00:22:19:27:54:c7", "nic mac address")
	kargs := flag.String("k", "", "boot kernel arguments")
	remove := flag.String("remove", "null", "remove systemt from cobbler")
	flag.Parse()

	if *remove != "null" {
		removeSystem(*remove)
	} else {
		newSystem(*machine, *profile, "managed by entitlement team", "testing", *kargs, *nicName, *nicMac)
		// rebootSystem(*machine)
	}
}
