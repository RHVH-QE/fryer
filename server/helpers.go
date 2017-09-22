package server

import (
	"github.com/dracher/fryer/model"
	"github.com/gin-gonic/gin"
)

func getQuery(c *gin.Context) *model.Query {
	r, _ := c.Get("q")
	return r.(*model.Query)
}
