package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/dracher/rhvhhelpers/model"
	"github.com/dracher/rhvhhelpers/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	db, err := gorm.Open("sqlite3", "./rhvhelper.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conn := model.Database{DB: db}
	conn.InitDB()

	app := server.InitApp(false, &conn)
	app.Run(":8000")
}
