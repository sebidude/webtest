package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func GetEnvOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) > 0 {
		return v
	}
	return defaultValue
}

func main() {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Any("/", anyHandler)
	r.Run(GetEnvOrDefault("LISTEN_ADDRESS", ":8080"))
}

func anyHandler(c *gin.Context) {
	c.String(200, "Hello World.")
}
