package handler

import (
	"fmt"
	"net/http"
	"workspace/chat/model"
	"workspace/chat/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type chatinfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func GetChats(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			username, _ := util.GetUsername(token)
			user := model.User{}
			db.Where("user_name = ?", username).Find(&user)
			pvusers := []model.User{}

			db.Table("users as us").
				Select("DISTINCT us.*").
				Joins("INNER join messages as msg on msg.sender_refer=us.id or msg.recever_refer=us.id").
				Where("msg.recever_refer = ? OR msg.sender_refer=?", user.ID, user.ID).
				Scan(&pvusers)
			chats := []chatinfo{}
			for _, u := range pvusers {
				chats = append(chats, chatinfo{Name: u.Name, Username: u.UserName})
			}
			fmt.Println(chats)
			c.IndentedJSON(http.StatusOK, chats)
		}
	}
}

func ViewPanel() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !util.IsLogedIn(c) {
			c.Redirect(http.StatusFound, "/view/login")
		}
		util.BackToLogin(c)
		cookie, _ := c.Cookie("token")
		username, _ := util.GetUsername(cookie)
		data := gin.H{
			"username": username,
		}
		c.HTML(http.StatusOK, "panel.html", data)
	}
}
