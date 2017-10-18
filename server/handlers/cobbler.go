package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dracher/fryer/utils/cobbler"
	"github.com/gin-gonic/gin"
)

const (
	// OK is
	OK = http.StatusOK
)

// SystemParams is
type SystemParams struct {
	Name    string
	Profile string
	Comment string
	Status  string
	Kargs   string
	Nic     []string
}

// CobblerHandler is
func CobblerHandler(c *gin.Context) {
	cb := cobbler.NewCobbler()
	var sp SystemParams
	if c.Request.Method == "POST" {
		if err := c.BindJSON(&sp); err != nil {
			log.Error(err.Error())
			c.String(http.StatusPreconditionFailed, "create system failed")
		} else {
			cb.NewSystem(sp.Name, sp.Profile, sp.Comment, sp.Status, sp.Kargs, sp.Nic)
			c.String(OK, "create system ok")
		}

	} else if c.Request.Method == "DELETE" {
		if err := c.BindJSON(&sp); err != nil {
			log.Error(err.Error())
			c.String(http.StatusPreconditionFailed, "remove system failed")
		} else {
			cb.RemoveSystem(sp.Name)
			c.String(OK, "remove system ok")
		}
	}
}
