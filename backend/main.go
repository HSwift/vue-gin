package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.StaticFile("/", "./dist/index.html")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")
	r.Static("/js", "./dist/js")
	r.Static("/css", "./dist/css")
	r.Static("/fonts", "./dist/fonts")

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
		code := c.PostForm("code")
		captchaid := c.PostForm("captchaid")
		var password string
		if username == "" || formPassword == "" || code == "" || captchaid == "" {
			c.JSON(http.StatusBadRequest, "Bad Request")
			return
		}
		if !captcha.VerifyString(captchaid, code) {
			c.JSON(http.StatusOK, gin.H{"code": -2, "msg": "验证码错误"})
			return
		}
		err := db.QueryRow("SELECT password from users WHERE username=?", username).Scan(&password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户名或密码错误"})
		} else {
			err = bcrypt.CompareHashAndPassword([]byte(password), []byte(formPassword))
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户名或密码错误"})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登录成功"})
			}
		}
	})

	r.POST("/api/register/:name", func(c *gin.Context) {
		username := c.Params.ByName("name")
		formPassword := c.PostForm("password")
		code := c.PostForm("code")
		captchaid := c.PostForm("captchaid")
		var count int64
		if username == "" || formPassword == "" || code == "" || captchaid == "" {
			c.JSON(http.StatusBadRequest, "Bad Request")
			return
		}
		if !captcha.VerifyString(captchaid, code) {
			c.JSON(http.StatusOK, gin.H{"code": -2, "msg": "验证码错误"})
			return
		}
		err := db.QueryRow("SELECT count(*) from users WHERE username=?", username).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		} else {
			if count > 0 {
				c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户已存在"})
			} else {
				hashPassword, err := bcrypt.GenerateFromPassword([]byte(formPassword), bcrypt.MinCost)
				if err != nil {
					c.JSON(http.StatusInternalServerError, "Internal Server Error")
					return
				}
				_, err = db.Exec("INSERT INTO users VALUES (NULL,?,?)", username, string(hashPassword))
				if err != nil {
					c.JSON(http.StatusInternalServerError, "Internal Server Error")
				} else {
					c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "注册成功"})
				}
			}
		}
	})

	r.GET("/captcha/:name", func(c *gin.Context) {
		captchaid := c.Param("name")
		err := captcha.WriteImage(c.Writer, captchaid, captcha.StdWidth, captcha.StdHeight)
		if err != nil {
			c.JSON(http.StatusNotFound, "File Not Found")
		}
	})

	r.GET("/captcha", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "获取成功", "captchaid": captcha.New()})
	})

	return r
}

func main() {

	db, err = sql.Open("mysql", "test:test1234@tcp(db:3306)/users?charset=utf8")
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
	r.Run(":8080")
}
