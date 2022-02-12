package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarathep/tutorial-backend/apis"
)

// Router to apis lisening
type Router struct {
	TutorialAPIs apis.TutorialHandler
}

// CORSMiddleware is Cross-Origin Resource Sharing Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func health(c *gin.Context) {

	type Resp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, Resp{
		Code:    "200",
		Message: "ok",
	})

}

// Route is setup router
func (router Router) Route() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/health", health)

	r.GET("/api/tutorials", router.TutorialAPIs.ReadTutorials)
	r.GET("/api/tutorials/:id", router.TutorialAPIs.ReadTutorial)
	r.POST("/api/tutorials", router.TutorialAPIs.CreateTutorial)
	r.PUT("/api/tutorials", router.TutorialAPIs.UpdateTutorial)
	r.PUT("/api/tutorials/:id", router.TutorialAPIs.UpdateTutorial)
	r.DELETE("/api/tutorials/:id", router.TutorialAPIs.DeleteTutorial)
	r.DELETE("/api/tutorials", router.TutorialAPIs.DeleteTutorials)

	return r
}
