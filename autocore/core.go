package autocore

import (
	"sync"
)

// Scheduler is
type Scheduler struct {
	MachinePool sync.Map
}

func (s *Scheduler) init() {
	
}