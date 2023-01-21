package handler

import (
	"fmt"
	"net/http"
	"workspace/chat/model"
	"workspace/chat/util"
	_ "workspace/chat/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MessageAndSenderinfo struct {
	Text       string
	Sendername string
}
type MessageRequest struct {
	Text         string `json:"text"`
	SenderUname  string `json:"senderuname"`
	ReceverUname string `json:"receveruname"`
}

func ViewPv(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !util.IsLogedIn(c) {
			c.Redirect(http.StatusFound, "/view/login")
		}
		token, err := c.Cookie("token")
		if err == nil {
			util.BackToLogin(c)
		}
		username, err := util.GetUsername(token)
		if err != nil {
			util.BackToLogin(c)
		}
		firstUserId := util.GetUserIDByUsername(*db, username)
		secondUserId := util.GetUserIDByUsername(*db, c.Param("username"))

		count := 0
		messages := []MessageAndSenderinfo{}
		db.Table("messages as msg").
			Select("u1.name as sendername ,msg.text").
			Joins("INNER join users as u1 on msg.sender_refer=u1.id ").
			Joins("INNER join users as u2 on msg.recever_refer=u2.id").
			Where("(msg.recever_refer = ? AND msg.sender_refer=?) OR (msg.recever_refer = ? AND msg.sender_refer=?)",
				firstUserId, secondUserId, secondUserId, firstUserId).
			Scan(&messages).Count(&count)
		data := gin.H{
			"pms":         messages,
			"count":       count,
			"username":    username,
			"recevername": c.Param("username"),
		}
		c.HTML(http.StatusOK, "pv.html", data)
	}
}

func CreateMessage(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !util.IsLogedIn(c) {
			c.Status(http.StatusBadRequest)
		}

		message := MessageRequest{}
		c.Bind(&message)

		token, err := c.Cookie("token")
		if err == nil {
			util.BackToLogin(c)
		}
		username, err := util.GetUsername(token)
		if message.SenderUname != username {
			c.Status(http.StatusBadRequest)
		}
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		fmt.Println(".................................")
		fmt.Println(message)
		fmt.Println(".................................")
		senderId := util.GetUserIDByUsername(*db, message.SenderUname)
		receverId := util.GetUserIDByUsername(*db, message.ReceverUname)
		err = db.Create(&model.Message{SenderRefer: uint(senderId),
			ReceverRefer: uint(receverId),
			Text:         message.Text}).Error
		if err == nil {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusBadRequest)
		}
	}
}
