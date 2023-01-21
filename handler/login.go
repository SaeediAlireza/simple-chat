package handler

import (
	"fmt"
	"net/http"
	"workspace/chat/model"
	"workspace/chat/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Logininfo struct {
	Username string `json:"username"`
	Pass     string `jsin:"pass"`
}

func Login(d *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		logininfo := Logininfo{}
		c.Bind(&logininfo)
		fmt.Println(logininfo)
		// fmt.Println(logininfo)
		count := 0
		err := d.Model(&model.User{}).
			Where("password = ? AND user_name = ?", logininfo.Pass, logininfo.Username).
			Count(&count).
			Error
		if err != nil {
			panic(err)
		}
		exists := count > 0

		if exists {
			var newCookie string
			newCookie = "NotSet"
			cookie, err := util.GenerateJWT(logininfo.Username)
			if err != nil {
				fmt.Println(err)
			} else {
				newCookie = cookie
			}
			fmt.Println("lxckjhguuuuuuuuuuuuuuuuuuuuuuuuuulx")
			c.SetCookie("token", newCookie, 360, "/", "localhost", false, true)
			c.Status(http.StatusOK)

		} else {
			c.Status(http.StatusFound)
		}

	}
}
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "NotSet", 360, "/", "localhost", false, true)
		c.Redirect(http.StatusFound, "/view/login")

	}
}
func ViewLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if util.IsLogedIn(c) {
			c.Redirect(http.StatusFound, "/view/panel")
		}
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
