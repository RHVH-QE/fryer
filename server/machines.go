package server

import "sync"

// MachinePool record current machine usage
type MachinePool struct {
	Pool map[string]bool
	lock *sync.RWMutex
}

// NewMachinePool is
func NewMachinePool() *MachinePool {
	return &MachinePool{
		Pool: make(map[string]bool),
		lock: &sync.RWMutex{},
	}
}

// CheckInUse is
func (m MachinePool) CheckInUse(name string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.Pool[name]
	return ok
}

// MarkMachine is
func (m *MachinePool) MarkMachine(name string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Pool[name] = true
}

// MachinesList is
func (m MachinePool) MachinesList() map[string]bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.Pool
}
