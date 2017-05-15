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
	im *service.IMService
)

// WebIMAPI main api endpoint for webim
func WebIMAPI(port string, dbs *service.DBService) {
	router := gin.Default()
	im = service.NewIMService(dbs)

	// use session
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("webim-session", store))
	// use cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://tinyurl.api.adolphlwq.xyz"},
		AllowMethods: []string{"*"},
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

	userAPI := router.Group("/api/v1/user")
	{
		userAPI.POST("/login", UserLogin)
		userAPI.POST("/register", UserRegister)
		userAPI.POST("/logout", LoginOut)
		userAPI.GET("/get", GetUserByName)
	}

	friendAPI := router.Group("/api/v1/friend")
	{
		friendAPI.POST("/add", AddFriend)
		friendAPI.GET("/list", ListFriend)
	}

	router.Run(port)
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

// HealthCheck return "health" if everything is OK
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "health"})
}
