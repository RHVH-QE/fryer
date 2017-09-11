package autocore

import (
	"github.com/looplab/fsm"
)

// Installation for manual provsion
var Installation = []fsm.EventDesc{
	{Name: "preJobCobbler", Src: []string{"jobFinished"}, Dst: "pxeReady"},
	{Name: "preJobBeaker", Src: []string{"pxeReady"}, Dst: "machineRebooted"},
	{Name: "waitUntilProvsionDone", Src: []string{"machineRebooted"}, Dst: "provisonDone"},
}

// AutoInstallation for auto install testing
var AutoInstallation = append(Installation, []fsm.EventDesc{
	{Name: "waitHostUp", Src: []string{"provisonDone"}, Dst: "hostAvailable"},
	{Name: "runTestCases", Src: []string{"hostAvailable"}, Dst: "testrunCompleted"},
	{Name: "handleLogs", Src: []string{"testrunCompleted"}, Dst: "logProcessed"},
	{Name: "postJobs", Src: []string{"logProcessed"}, Dst: "jobFinished"},
}...)
