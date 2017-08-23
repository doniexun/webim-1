package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	session sessions.Session
)

func builfEngine(appService *ServiceProvider) *gin.Engine {
	// use session
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 30 * 60, //30mins
		Path:   "/",
	})

	router := gin.Default()
	router.Use(sessions.Sessions("webim-session", store))
	// use cors
	router.Use(cors.New(cors.Config{
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

	router.GET("/health", WrapeService(appService, HealthCheck))
	userAPI := router.Group("/api/v1/user")
	{
		userAPI.POST("/register", WrapeService(appService, UserRegister))
		userAPI.POST("/login", WrapeService(appService, UserLogin))
		userAPI.GET("/list", WrapeService(appService, UserList))
	}
	contactAPI := router.Group("/api/v1/contact")
	{
		contactAPI.POST("/add", WrapeService(appService, ContactAdd))
		contactAPI.DELETE("/delete", WrapeService(appService, ContactDelete))
	}

	return router
}

// Start start webim server
func Start(addr string, appService *ServiceProvider) {
	router := builfEngine(appService)
	router.Run(addr)
}

// RequestHandler alias
type RequestHandler func(*gin.Context, *ServiceProvider)

// WrapeService wrape instance needed by gin.Engine
func WrapeService(appService *ServiceProvider, handler RequestHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, appService)
	}
}
