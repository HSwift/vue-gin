package main

import (
	"database/sql"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/api/user/:name", func(c *gin.Context) {
		username := c.Params.ByName("name")
		var password string
		err := db.QueryRow("SELECT password from users WHERE username=?", username).Scan(&password)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{"username": username, "password": nil})
		} else {
			c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
		}
	})

	r.POST("/api/login/:name", func(c *gin.Context) {
		username := c.Params.ByName("name")
		formPassword := c.PostForm("password")
		var password string
		err := db.QueryRow("SELECT password from users WHERE username=?", username).Scan(&password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户名或密码错误"})
		} else {
			if formPassword == password {
				c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登录成功"})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户名或密码错误"})
			}
		}
	})

	r.POST("/api/register/:name", func(c *gin.Context) {
		username := c.Params.ByName("name")
		formPassword := c.PostForm("password")
		var count int64
		err := db.QueryRow("SELECT count(*) from users WHERE username=?", username).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": -2, "msg": "数据库错误"})
		} else {
			if count > 0 {
				c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户已存在"})
			} else {
				_, err := db.Exec("INSERT INTO users VALUES (?,?)", username, formPassword)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": -2, "msg": "数据库错误"})
				} else {
					c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "注册成功"})
				}
			}
		}
	})

	return r
}

func main() {

	db, err = sql.Open("mysql", "root:your_password@tcp(127.0.0.1:3306)/users?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8090")
}
