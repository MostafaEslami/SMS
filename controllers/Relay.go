package controllers

import (
	"ECommerce/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
func MakeRequest(mobile string, code string) string {

	//request := fmt.Sprintf("http://5m5.ir/send_via_get/send_sms_by_pattern.php?username=khadamati1400&password=WbUqSBo&receiver_number=%s&pattern_id=%s&pattern_params[]=%s", mobile, pattern, code)
	request := fmt.Sprintf("http://robots.rahco.ir/api/proxy/send?username=khadamati1400&password=WbUqSBo&receiver_number=%s&pattern_id=38&pattern_params[]=%s&token=rVW9HmLH41RjA5PywpuHGdfXODzbQo", mobile, code)
	return request
}

func SendSMS(c *gin.Context) {

	mobile := c.PostForm("receiver_number")
	code := c.PostForm("code")
	if mobile == "" || code == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("some params is required!"))
	}

	request := MakeRequest(mobile, code)

	fmt.Println("request:", request)
	resp, err := http.Get(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Log(3, "%s", err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	//c.JSON(http.StatusBadRequest, string(request))
	//return
	s := fmt.Sprintf("%s", body)
	cdr := utility.CDR{Number: mobile, Code: code, MessageId: s}
	LogCDR(cdr)
	Log(2, cdr.Log())
	c.JSON(http.StatusOK, gin.H{
		"data": s,
	})
}
