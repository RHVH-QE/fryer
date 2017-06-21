package beaker

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

var cmds = map[string]string{
	"reboot": "bkr system-power --action reboot %s",
	"on":     "bkr system-power --action on %s",
	"off":    "bkr system-power --action off %s",
}

// Beaker wraps beaker-cli commands
type Beaker struct {
	SystemName string
}

// NewBeaker is
func NewBeaker(bkrName string) *Beaker {
	return &Beaker{bkrName}
}

// Reboot is
func (b Beaker) Reboot() []byte {
	out, err := exec.Command("bkr", "system-power", "--action", "reboot",
		b.SystemName).Output()

	if err != nil {
		log.Error(err)
	}
	log.Info(out)
	return out
}

// PowerOn is
func (b Beaker) PowerOn() {
	out, err := exec.Command("bkr", "system-power", "--action", "on",
		b.SystemName).Output()

	if err != nil {
		log.Error(err)
	}
	log.Info(out)
}

// PowerOff is
func (b Beaker) PowerOff() {
	out, err := exec.Command("bkr", "system-power", "--action", "off",
		b.SystemName).Output()

	if err != nil {
		log.Error(err)
	}
	log.Info(out)
}

// TODO setup beaker client if not exists
func (b Beaker) SetupBeakerClient() {
	panic("Not Implemented")
}
