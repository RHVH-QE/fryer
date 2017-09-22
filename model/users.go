package model

import "time"

// User is
type User struct {
	ID        int    `storm:"id,increment"`
	Group     string `storm:"index"`
	Email     string `storm:"unique"`
	KrbID     string `storm:"unique,index"`
	Password  string
	CreatedAt time.Time `storm:"index"`
	Admin     bool
}
