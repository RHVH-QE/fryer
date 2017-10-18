package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/dracher/fryer/helper"
)

// ConfigParamsHandler is
func ConfigParamsHandler(c *gin.Context) {
	name := c.Param("name")
	q := helper.GetQuery(c)
	switch name {
	case "common":
		ret, _ := q.CommonParams()
		c.JSON(OK, ret)
	case "hosts":
		ret, _ := q.Hosts()
		c.JSON(OK, ret)
	case "tiers":
		ret, _ := q.AutoTestTiers()
		c.JSON(OK, ret)
	case "casemap":
		ret, _ := q.AutoTestCaseMap()
		c.JSON(OK, ret)
	default:
		c.JSON(OK, nil)
	}
}
