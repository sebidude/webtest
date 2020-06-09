package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	hostname    string
	greeting    string
	contentdir  string
	contentfile string
	filecontent []byte
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
	contentfile = GetEnvOrDefault("CONTENTFILE", "/content/content.txt")
	contentdir = GetEnvOrDefault("CONTENTDIR", "/content")
	filecontent, err = ioutil.ReadFile(contentfile)
	if err != nil {
		fmt.Printf("Cannot read content from %s: %s", contentfile, err.Error())
		os.Exit(1)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Any("/", helloHandler)
	r.GET("/filecontent", contentHandler)
	r.StaticFS("/contentdir", gin.Dir(contentdir, true))
	r.Run(GetEnvOrDefault("LISTEN_ADDRESS", ":8080"))
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hello World from %s (%s)", hostname, greeting)
}

func contentHandler(c *gin.Context) {
	c.String(200, string(filecontent))
}
