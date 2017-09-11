package main

import (
	"github.com/dracher/fryer/model"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	model.ImportData()
}
