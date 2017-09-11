package model

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/msgpack"
)

func ImportData() {

	db, _ := storm.Open("fryer.conf.db", storm.Codec(msgpack.Codec))

	var caseMap []TestCaseMap

	err := db.All(&caseMap)
	fmt.Println(err)

	fmt.Println(caseMap)
}
