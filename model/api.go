package model

import (
	"crypto/sha1"
	"fmt"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/msgpack"
	log "github.com/sirupsen/logrus"
)

func openDB() *storm.DB {
	db, err := storm.Open("fryer.conf.db", storm.Codec(msgpack.Codec))
	if err != nil {
		log.Panic(err)
	}
	return db
}

// InitDB is
func InitDB() {
	db := openDB()
	for _, user := range users {
		user.Password = fmt.Sprintf("%x", sha1.Sum([]byte(user.Password)))
		db.Save(&user)
	}
	db.Close()
}

// Query holds all data from db
type Query struct {
	DB *storm.DB
}

// NewQuery is
func NewQuery() *Query {
	return &Query{
		DB: openDB(),
	}
}

// User is
func (q Query) User(field, val string) (User, error) {
	var user User
	err := q.DB.One(field, val, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Users is
func (q Query) Users() ([]User, error) {
	var users []User
	err := q.DB.All(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}

// CheckUser is
func (q Query) CheckUser(krbID, password string) (string, bool) {
	user, err := q.User("KrbID", krbID)
	if err != nil {
		log.Error(err)
		return krbID, false
	}
	if user.Password != fmt.Sprintf("%x", sha1.Sum([]byte(password))) {
		return krbID, false
	}
	return krbID, true
}

// CommonParams is
func (q Query) CommonParams() (CommonParams, error) {
	var params CommonParams
	err := q.DB.One("ID", "MainConfig", &params)
	if err != nil {
		return params, err
	}
	return params, nil
}

// Host is
func (q Query) Host(filed, val string) (Host, error) {
	var host Host
	err := q.DB.One(filed, val, &host)
	if err != nil {
		return host, err
	}
	return host, nil
}

// Hosts is
func (q Query) Hosts() ([]Host, error) {
	var hosts []Host
	err := q.DB.All(&hosts)
	if err != nil {
		return hosts, err
	}
	return hosts, nil
}

// AutoTestTiers is
func (q Query) AutoTestTiers() (AutoTestTiers, error) {
	var tiers AutoTestTiers
	err := q.DB.One("ID", "tiers", &tiers)
	if err != nil {
		return tiers, err
	}
	return tiers, nil
}

// AutoTestCaseMap is
func (q Query) AutoTestCaseMap() ([]TestCaseMap, error) {
	var caseMap []TestCaseMap
	err := q.DB.All(&caseMap)
	if err != nil {
		return caseMap, err
	}
	return caseMap, nil
}

// Close is
func (q Query) Close() {
	q.DB.Close()
}
