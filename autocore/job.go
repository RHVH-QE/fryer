package autocore

import (
	"github.com/looplab/fsm"
	log "github.com/sirupsen/logrus"

	"github.com/dracher/rhvhhelpers/utils/beaker"
	"github.com/dracher/rhvhhelpers/utils/cobbler"
)

// NewJob returns a job with unique name, noramlly is beaker name
func NewJob(jobType string) *fsm.FSM {
	return fsm.NewFSM(
		"jobFinished",
		Installation,
		fsm.Callbacks{
			"leave_jobFinished": func(e *fsm.Event) {
				cb := cobbler.NewCobbler()
				if len(e.Args) != 6 {
					log.Fatalf("event %s args must have 6 elements", e.Event)
				}
				cb.NewSystem(
					e.Args[0].(string),
					e.Args[1].(string),
					e.Args[2].(string),
					e.Args[3].(string),
					e.Args[4].(string),
					e.Args[5].([]string))
			},
			"leave_pxeReady": func(e *fsm.Event) {
				bkr := beaker.NewBeaker(e.Args[0].(string))
				out, err := bkr.Reboot()
				if err != nil {
					log.Error(err, out)
					e.Cancel()
				}
			},
		},
	)
}
