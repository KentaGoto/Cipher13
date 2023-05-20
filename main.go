package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ROT13 conversion
func rot13(s string) string {
	rot13 := func(r rune) rune {
		switch {
		case 'a' <= r && r <= 'z':
			return 'a' + (r-'a'+13)%26
		case 'A' <= r && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		default:
			return r
		}
	}
	return strings.Map(rot13, s)
}

func main() {
	router := gin.Default()

	// API Routing
	router.GET("/api/rot13", func(c *gin.Context) {
		// Get the original string from the query parameter.
		original := c.Query("s")
		// ROT13 conversion of the original string.
		encrypted := rot13(original)
		// Return results as JSON.
		c.JSON(200, gin.H{
			"original": original,
			"rot13":    encrypted,
		})
	})

	// Serving Static Files.
	router.Static("/static", "./my-app/build/static")
	router.GET("/", func(c *gin.Context) {
		c.File("./my-app/build/index.html")
	})

	router.Run(":8080")
}
