package autocore

import (
	"fmt"
)

func CheckHostAvailable(bkrName, userName string) error {
	Scheduler.Lock.Lock()
	defer Scheduler.Lock.Unlock()

	if !Scheduler.Pool[bkrName].Available {
		return fmt.Errorf("host %s already in use", bkrName)
	}
	Scheduler.SetAvailable(bkrName, userName, false)
	return nil
}
