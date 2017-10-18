package autocore

import (
	"sync"

	"github.com/dracher/fryer/helper"
	"github.com/dracher/fryer/model"
	"github.com/looplab/fsm"
)

var log = helper.GetZapLogger()

// Attrs is
type Attrs struct {
	ReseverdBy string
	Available  bool
	Jobs       []*fsm.FSM
	SessionID  string
}

type scheduler struct {
	Lock *sync.RWMutex
	Pool map[string]Attrs
}

// Scheduler is
var Scheduler = &scheduler{
	Lock: &sync.RWMutex{},
	Pool: make(map[string]Attrs),
}

func (s *scheduler) SetAvailable(bkrName, userName string, b bool) {
	s.Pool[bkrName] = Attrs{
		ReseverdBy: userName,
		Available:  b,
		Jobs:       s.Pool[bkrName].Jobs,
		SessionID:  helper.RandStringBytesMaskImprSrc(12),
	}
}

// InitScheduler is
func InitScheduler(hosts []model.Host) {
	for _, host := range hosts {
		Scheduler.Pool[host.BeakerName] = Attrs{
			ReseverdBy: "nobody",
			Available:  true,
		}
	}
	log.Infof("Init scheduler with %s finished", hosts)
}
