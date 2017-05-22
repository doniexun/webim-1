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
	router.GET("/wscount", WSCount)

	userAPI := router.Group("/api/v1/user")
	{
		userAPI.POST("/login", UserLogin)
		userAPI.POST("/register", UserRegister)
		userAPI.POST("/logout", LogOut)
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

// HealthCheck return "health" if everything is OK
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "health"})
}

// WSCount return count of websockets
func WSCount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   im.UserWSMap.Length()})
}
