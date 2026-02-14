package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const serverAddr = ":8080"

type rot13Response struct {
	Original string `json:"original"`
	Rot13    string `json:"rot13"`
}

// ROT13 conversion
func rot13(s string) string {
	rotate := func(r rune) rune {
		switch {
		case 'a' <= r && r <= 'z':
			return 'a' + (r-'a'+13)%26
		case 'A' <= r && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		default:
			return r
		}
	}

	return strings.Map(rotate, s)
}

func handleRot13(c *gin.Context) {
	original := c.Query("s")
	response := rot13Response{
		Original: original,
		Rot13:    rot13(original),
	}

	c.JSON(http.StatusOK, response)
}

func registerRoutes(router *gin.Engine) {
	router.GET("/api/rot13", handleRot13)
	router.Static("/static", "./my-app/build/static")
	router.GET("/", func(c *gin.Context) {
		c.File("./my-app/build/index.html")
	})
}

func newRouter() *gin.Engine {
	router := gin.Default()
	registerRoutes(router)
	return router
}

func main() {
	router := newRouter()
	router.Run(serverAddr)
}
