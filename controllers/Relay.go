package controllers

import (
	"ECommerce/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "os/exec"
	_ "runtime"
	"time"
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
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

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

func DoneAsync(mobile string, code string, generatedMsgId string) chan int {
	r := make(chan int)
	go func() {
		request := MakeRequest(mobile, code)

		//fmt.Println("request : ", request)
		resp, err1 := http.Get(request)
		if resp == nil || resp.Body == nil || err1 != nil {
			cdr := utility.CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: "FAILED"}
			utility.LogCDR(cdr)
			utility.Log("ERROR", cdr.Log())
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			cdr := utility.CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: "FAILED"}
			utility.LogCDR(cdr)
			utility.Log("ERROR", cdr.Log())
			return
		}

		s := fmt.Sprintf("%s", body)
		cdr := utility.CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: s}
		utility.LogCDR(cdr)
		utility.Log("INFO", cdr.Log())

	}()
	return r
}

func SendSMS(c *gin.Context) {
	if utility.HasCredit() == false {
		//utility.Log("WARNING", "credit error")
		c.JSON(http.StatusBadRequest, CreateErrorMessage("credit error"))
		return
	}
	//c.JSON(http.StatusBadRequest, CreateErrorMessage("has credit"))
	//return
	mobile := c.PostForm("receiver_number")
	code := c.PostForm("code")
	if mobile == "" || code == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("some params is required!"))
		return
	}

	utility.DecreaseCreditAsync()

	rand.Seed(time.Now().UnixNano())
	randomNum := random(100000000000, 200000000000)
	generatedMessageId := fmt.Sprintf("%d", randomNum)
	DoneAsync(mobile, code, generatedMessageId)
	c.JSON(http.StatusOK, gin.H{
		"data": generatedMessageId,
	})
}
