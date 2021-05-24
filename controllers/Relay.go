package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "os/exec"
	_ "runtime"
)

type Instance struct {
	Id             string
	URL            string
	Secret         string
	Enable         bool
	Load           string
	LoadMultiplier string
	Online         bool
}

//func CreateSuccessMessage(message string) gin.H {
//	ret := gin.H{
//		"success": true,
//		"message": message,
//	}
//	return ret
//}
//
func CreateErrorMessage(message string) gin.H {
	ret := gin.H{
		"success": false,
		"message": message,
	}
	return ret
}

func RelayRoutes(router *gin.RouterGroup) {
	{
		router.POST("/send", SendSMS)
		//router.GET("/status", SendStatus)
	}
}
func MakeRequest(mobile string, code string, pattern string) string {

	request := fmt.Sprintf("http://5m5.ir/send_via_get/send_sms_by_pattern.php?username=khadamati1400&password=WbUqSBo&receiver_number=%s&pattern_id=%s&pattern_params[]=%s", mobile, pattern, code)
	return request
}

func SendSMS(c *gin.Context) {

	mobile := c.PostForm("receiver_number")
	code := c.PostForm("code")
	if mobile == "" || code == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("some params is required!"))
	}

	request := MakeRequest(mobile, code, "38")
	fmt.Println("request:", request)
	resp, err := http.Get(request)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
	}
	//c.JSON(http.StatusBadRequest, string(request))
	//return
	c.JSON(http.StatusOK, resp)
}
