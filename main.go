package main

import (
	"ECommerce/controllers"
	"ECommerce/utility"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"os"
	"time"
)

func main() {

	controllers.Initialize()

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	goGonicEngine := gin.New()
	goGonicEngine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	goGonicEngine.Use(gin.Recovery())
	goGonicEngine.Use(cors.Default())

	// goGonicEngine.Use(middlewares.Cors())

	apiRouteGroup := goGonicEngine.Group("/api")

	controllers.RelayRoutes(apiRouteGroup.Group("/relay"))

	cdr := utility.CDR{From: "sasa", To: "sasasa"}
	controllers.LogCDR(cdr)
	controllers.Log(245, "dsadsadsadsa")

	goGonicEngine.Run(":8080") // listen and serve on 0.0.0.0:8080

}
