package handlers

import (
	"fmt"
	"net/http"

	"github.com/dracher/fryer/helper"
	"github.com/dracher/fryer/model"
	"github.com/gin-gonic/gin"
	// "github.com/dracher/fryer/model"
)

// AddHostHandler is
func AddHostHandler(c *gin.Context) {
	var host model.Host
	q := helper.GetQuery(c)
	c.BindJSON(&host)
	err := q.DB.Save(&host)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ret": err})
	} else {
		c.JSON(http.StatusOK, host)
	}
}

// UpdateHandler is
func UpdateHandler(c *gin.Context) {
	var host model.Host
	q := helper.GetQuery(c)
	c.BindJSON(&host)
	fmt.Println(host)
	err := q.DB.Update(&host)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ret": err})
	} else {
		c.JSON(http.StatusOK, host)
	}
}
