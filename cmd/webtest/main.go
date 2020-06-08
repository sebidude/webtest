package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	hostname string
	greeting string
)

func GetEnvOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) > 0 {
		return v
	}
	return defaultValue
}

func main() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		fmt.Println("Cannot get hostname!")
		os.Exit(1)
	}

	greeting = GetEnvOrDefault("GREETING", "simple webtest")

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Any("/", anyHandler)
	r.Run(GetEnvOrDefault("LISTEN_ADDRESS", ":8080"))
}

func anyHandler(c *gin.Context) {
	c.String(200, "Hello World from %s (%s)", hostname, greeting)
}
