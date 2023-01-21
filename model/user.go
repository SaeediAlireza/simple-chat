package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	UserName string
	Password string
	Gender   bool
	Birthday time.Time
}

type Pv struct {
	gorm.Model
	Owner        User `gorm:"foreignkey:WriterRefer"`
	OwnerRefer   uint
	Contact      User `gorm:"foreignkey:WriterRefer"`
	ContactRefer uint
}

type Message struct {
	gorm.Model
	Sender       User `gorm:"foreignkey:WriterRefer"`
	SenderRefer  uint
	Recever      User `gorm:"foreignkey:WriterRefer"`
	ReceverRefer uint
	Text         string
	Seen         bool
}

type Reply struct {
	gorm.Model
	ReplyTo      Message `gorm:"foreignkey:ReplyToRefer"`
	ReplyToRefer uint
	Reply        Message `gorm:"foreignkey:ReplyRefer"`
	ReplyRefer   uint
}
