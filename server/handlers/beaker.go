package handlers

import (
	"fmt"

	"github.com/dracher/fryer/utils/beaker"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// BeakerHandler is
func BeakerHandler(c *gin.Context) {
	name := c.Param("bkrname")
	action := c.Param("action")
	bkr := beaker.Beaker{SystemName: name}

	switch action {
	case "reboot":
		log.Infof("%s %s", action, name)
		bkr.Reboot()
	case "on":
		log.Infof("%s %s", action, name)
		bkr.PowerOn()
	case "off":
		log.Infof("%s %s", action, name)
		bkr.PowerOff()
	default:
		log.Infof("no action performed on %s", name)
	}
	c.String(OK, fmt.Sprintf("%s %s is complete", name, action))
}
