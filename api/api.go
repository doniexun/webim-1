package api

import (
	"net/http"
	"time"

	"webim/db"
	"webim/service"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	us *service.UserService
)

// WebIMAPI main api endpoint for webim
func WebIMAPI(port string, dbs *db.DBService) {
	router := gin.Default()
	us = service.NewUserService(dbs)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://tinyurl.api.adolphlwq.xyz"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		//AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			logrus.Info("origin is ", origin)
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/health", HealthCheck)

	userAPI := router.Group("/api/v1/user")
	{
		userAPI.GET("/login", UserLogin)
		userAPI.POST("/register", UserRegister)
	}

	router.Run(port)
}

// UserRegister handle user register
func UserRegister(c *gin.Context) {
	var user db.User
	c.BindJSON(&user)

	err := us.UserRegister(user.Username, user.Password)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
			"data": "register successfully, please login."})
	} else {
		logrus.Warn("register user info error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
			"data": "register error, please try again."})
	}
}

// UserLogin handle user login
func UserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "userlogin"})
}

// HealthCheck return "health" if everything is OK
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "health"})
}
