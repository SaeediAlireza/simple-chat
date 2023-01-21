package util

import (
	"fmt"
	"net/http"
	"time"
	"workspace/chat/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// session Connecting to Database
var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:@alireza0918@(localhost:3306)/goapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() gorm.DB {
	return *db
}

// end session

// session Genarate and Author Json Web Token
var mySigningKey = []byte("mysupersecretphrase")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenstring, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something is not OK: %s", err.Error())
		return "", err
	}
	return tokenstring, err
}
func IsAuthorizes(tok string) (bool, error) {
	token, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("something is wrong!")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false, fmt.Errorf("Something is not OK: %s", err.Error())
	}
	if token.Valid {
		sdjf := token.Claims.(jwt.MapClaims)
		fmt.Println(sdjf["user"])
		return true, nil
	} else {
		return false, nil
	}

}

// end session
func GetUsername(tok string) (string, error) {
	token, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("something is wrong!")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("Something is not OK: %s", err.Error())
	}
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		uerame := claim["user"]
		return uerame.(string), nil
	} else {
		return "", nil
	}

}

// session is login
func IsLogedIn(c *gin.Context) bool {
	cookie, err := c.Cookie("token")
	if isauthorizes, er := IsAuthorizes(cookie); er == nil && isauthorizes && err == nil {
		return true
	} else {
		return false
	}
}
func BackToLogin(c *gin.Context) {
	if !IsLogedIn(c) {
		c.Redirect(http.StatusFound, "/view/login")
	}
}

func GetUserIDByUsername(db gorm.DB, username string) int {
	user := model.User{UserName: username}
	db.Where("user_name = ?", username).First(&user)
	return int(user.ID)
}
