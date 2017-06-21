package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User is
type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Username  string     `gorm:"primary_key"`
	Password  string
	Regular   bool
}

// Database is
type Database struct {
	DB *gorm.DB
}

// InitDB is
func (d Database) InitDB() {
	d.DB.AutoMigrate(&User{})
	d.insertInitData()
}

func (d Database) insertInitData() {
	users := []User{
		User{
			Username: "dracher",
			Password: "rhvher",
			Regular:  true,
		},
		User{
			Username: "tester",
			Password: "tester",
			Regular:  false,
		},
	}
	for _, user := range users {
		d.DB.FirstOrCreate(&user)
	}
}

// CheckUser is
func (d Database) CheckUser(username, password string) (res User) {
	d.DB.Where("username = ? AND password = ?", username, password).Find(&res)
	return
}

// FindUser is
func (d Database) FindUser(username string) (res User) {
	d.DB.Where("username = ?", username).First(&res)
	return
}
