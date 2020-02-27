package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var router = gin.Default()

func StartApplication() {
	urlMapper()
	router.Run(fmt.Sprintf(":%s", os.Getenv("SERVERPORT")))
}
