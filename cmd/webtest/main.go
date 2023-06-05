package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	hostname     string
	greeting     string
	contentdir   string
	contentfile  string
	filecontent  []byte
	readycounter = 0
	alivecounter = 0
	failcounter  = 0
	latefail     int
	latesuccess  int
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
	latefail = 20 - rand.Intn(20)
	latesuccess = 20 - rand.Intn(10)
	fmt.Printf("latefail: %d\n", latefail)
	fmt.Printf("latesuccess: %d\n", latesuccess)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Any("/", helloHandler)
	r.GET("/filecontent", contentHandler)
	r.StaticFS("/contentdir", gin.Dir(contentdir, true))
	r.GET("/ready", readyProbe)
	r.GET("/alive", aliveProbe)
	r.GET("/fail", failProbe)
	r.GET("/faillate", faillateProbe)
	r.GET("/readylate", readylateProbe)
	r.Run(GetEnvOrDefault("LISTEN_ADDRESS", ":8080"))
}

func readyProbe(c *gin.Context) {
	if readycounter <= 2 {
		c.String(500, "Not ready yet.")
		readycounter = readycounter + 1
		return
	}
	c.String(200, "Ready.")
}

func aliveProbe(c *gin.Context) {
	if alivecounter <= 1 {
		c.String(500, "Not alive yet.")
		alivecounter = alivecounter + 1
		return
	}
	c.String(200, "Alive.")
}

func faillateProbe(c *gin.Context) {
	if failcounter < latesuccess {
		c.String(200, "Happy.")
		failcounter = failcounter + 1
		return
	}
	failcounter = failcounter + 1
	if failcounter > latefail {
		failcounter = 0
	}
	c.String(500, "Unhappy.")
}

func failProbe(c *gin.Context) {
	c.String(500, "Unhappy.")
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hello World from %s (%s)", hostname, greeting)
}

func contentHandler(c *gin.Context) {
	c.String(200, string(filecontent))
}
