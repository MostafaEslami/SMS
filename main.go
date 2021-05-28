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
	"strconv"
	"time"
)

func main() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	port := os.Getenv("PORT")
	utility.Initialize()

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
	utility.IniitalizeCredit()
	if len(os.Args) > 1 {
		creditInput := os.Args[1]
		creditEnter, _ := strconv.Atoi(creditInput)
		utility.SetCredit(creditEnter)
	}

	utility.Log("INFO", "credit is ", utility.GetCredit())

	apiRouteGroup := goGonicEngine.Group("/api")
	controllers.SetPattern(os.Getenv("PATTERN"))
	controllers.RelayRoutes(apiRouteGroup.Group("/sms"))

	//cdr := utility.CDR{From: "sasa", To: "sasasa"}
	//controllers.LogCDR(cdr)
	//controllers.Log(245, "dsadsadsadsa")

	goGonicEngine.Run(port) // listen and serve on 0.0.0.0:8080

}
