package main

import (
	// _ "workspace/chat/handler"
	"fmt"
	"workspace/chat/model"
	"workspace/chat/util"
)

type MessageAndSenderinfo struct {
	Text       string
	Sendername string
}

// Database Migration:
func main() {
	util.Connect()
	d := util.GetDB()
	d.AutoMigrate(&model.User{},
		// &model.Pv{},
		&model.Message{},
		&model.Reply{})

	// ms := model.Message{Text: "hi zahra", SenderRefer: 1, ReceverRefer: 2, Seen: false}
	// d.Create(&ms)
	// ms = model.Message{Text: "hi alireza", SenderRefer: 2, ReceverRefer: 1, Seen: false}
	// d.Create(&ms)

	messages := []MessageAndSenderinfo{}
	d.Table("messages as msg").
		Select("u1.name as sendername ,msg.text").
		Joins("INNER join users as u1 on msg.sender_refer=u1.id ").
		Joins("INNER join users as u2 on msg.recever_refer=u2.id").
		Where("msg.recever_refer = ? OR msg.sender_refer=? OR msg.recever_refer = ? OR msg.sender_refer=?",
			1, 1, 2, 2).
		Scan(&messages)
	fmt.Println(messages)
	defer d.Close()
}
