package api

import (
	"net/http"
	"time"

	"github.com/adolphlwq/webim/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	im      *service.IMService
	session sessions.Session
)

// WebIMAPI main api endpoint for webim
func WebIMAPI(serviceUrl string, dbs *service.DBService) {
	router := gin.Default()
	im = service.NewIMService(dbs)

	// use session
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 30 * 60, //30mins
	})
	router.Use(sessions.Sessions("webim-session", store))
	// use cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://tinyurl.api.adolphlwq.xyz"},
		AllowMethods: []string{"PUT", "GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		//AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/health", HealthCheck)
	router.GET("/incr", Incr)
	router.GET("/incr2", Incr2)

	userAPI := router.Group("/api/v1/user")
	{
		userAPI.POST("/login", UserLogin)
		userAPI.POST("/register", UserRegister)
		userAPI.POST("/logout", LoginOut)
		userAPI.GET("/get", GetUserByName)
	}

	friendAPI := router.Group("/api/v1/friend")
	{
		friendAPI.POST("/add", AddFriendRelationship)
		friendAPI.GET("/list", ListFriendRelationship)
		friendAPI.PUT("/delete", DeleteFriendRelationship)
	}

	messageAPI := router.Group("/api/v1/message")
	{
		messageAPI.GET("/unread", GetUnreadMsg)
		// add username(id) to path is inspired
		// my deep thinking and https://github.com/gin-gonic/gin/issues/461
		messageAPI.GET("/ws/:username", WSMsgHandler)
	}
	router.Run(serviceUrl)
}

// Incr
func Incr(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	c.JSON(200, gin.H{"count": count})
}

// Incr2
func Incr2(c *gin.Context) {
	var count int
	session := sessions.Default(c)
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
	}
	c.JSON(200, gin.H{"count": count})
}

// HealthCheck return "health" if everything is OK
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "health"})
}
