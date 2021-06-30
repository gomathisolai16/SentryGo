package main

import (
	"fmt"
	"log"

	sentry "github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	Routes(router)
	router.Run("0.0.0.0" + ":" + "7050")
}

//Routes method used for routing
func Routes(router *gin.Engine) {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "YOUR_PUBLIC_DSN",
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	router.Use(gin.Recovery())
	router.Use(sentrygin.New(sentrygin.Options{}))
	router.GET("/api/getList", GetList)
}

//List structure definition for getting list
type List struct {
	Name string `json:"name,omitempty"`
}

//GetList method handles to get the list of name
func GetList(c *gin.Context) {
	var list []List
	var listData List
	listData.Name = "Test Sentry"
	list = append(list, listData)
	fmt.Println("list", list[2]) // here am trying get the 2 nd element from array , but array contains only one element. So this error will capture in sentry
	c.JSON(200, gin.H{"message": "", "status": "Success", "code": 200, "data": list})
}
