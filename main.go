package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	port := os.Getenv("PORT")

	Initialize()

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
	IniitalizeCredit()
	if len(os.Args) > 1 {
		creditInput := os.Args[1]
		creditEnter, _ := strconv.Atoi(creditInput)
		SetCredit(creditEnter)
	}

	Log("INFO", "credit is ", GetCredit())

	apiRouteGroup := goGonicEngine.Group("/api")
	SetPattern(os.Getenv("PATTERN"))
	CheckRedis(os.Getenv("REDIS"))

	RelayRoutes(apiRouteGroup.Group("/sms"))

	//cdr := utility.CDR{From: "sasa", To: "sasasa"}
	//controllers.LogCDR(cdr)
	//controllers.Log(245, "dsadsadsadsa")

	goGonicEngine.Run(port) // listen and serve on 0.0.0.0:8080

}
