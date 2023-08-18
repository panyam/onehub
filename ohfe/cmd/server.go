package main

import (
	_ "html/template"
	"net/http"

	_ "github.com/Masterminds/sprig/v3"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Template interface {
	Serve(input map[string]interface{}, request *http.Request, writer *http.ResponseWriter)
}

func GetTemplate(name string) Template {
}

// Implement this
const HomePageTemplate = GetTemplate("HomePage")

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("onehub_sessions", store))
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
