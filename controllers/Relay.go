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
	"strconv"
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

var Pattern = "38"

var client = http.Client{
	Timeout: 10 * time.Second,
}

var mBlockList = make(map[string]bool)

func InitializeBlockList() {
	mBlockList["989132957573"] = true
	mBlockList["989128181275"] = true
	mBlockList["09128181275"] = true
	mBlockList["989128172212"] = true
	mBlockList["09128172212"] = true
	mBlockList["989121554134"] = true
	mBlockList["09121554134"] = true
	mBlockList["989124487370"] = true
	mBlockList["09124487370"] = true
	mBlockList["989131260622"] = true
	mBlockList["09131260622"] = true
	mBlockList["989122510039"] = true
	mBlockList["09122510039"] = true
	mBlockList["989132677411"] = true
	mBlockList["09132677411"] = true
	mBlockList["989132125459"] = true
	mBlockList["09132125459"] = true
	mBlockList["989123233064"] = true
	mBlockList["09123233064"] = true
	mBlockList["989123343946"] = true
	mBlockList["09123343946"] = true
	mBlockList["989131252502"] = true
	mBlockList["09131252502"] = true
	mBlockList["989370051430"] = true
	mBlockList["09370051430"] = true
	mBlockList["989111196076"] = true
	mBlockList["09111196076"] = true
	mBlockList["989121153684"] = true
	mBlockList["09121153684"] = true
	mBlockList["989121231333"] = true
	mBlockList["09121231333"] = true
	mBlockList["989121430087"] = true
	mBlockList["09121430087"] = true
	mBlockList["989121544643"] = true
	mBlockList["09121544643"] = true
	mBlockList["989121829499"] = true
	mBlockList["09121829499"] = true
	mBlockList["989121908612"] = true
	mBlockList["09121908612"] = true
	mBlockList["989121973544"] = true
	mBlockList["09121973544"] = true
	mBlockList["989122088891"] = true
	mBlockList["09122088891"] = true
	mBlockList["989122405575"] = true
	mBlockList["09122405575"] = true
	mBlockList["989122440411"] = true
	mBlockList["09122440411"] = true
	mBlockList["989122447108"] = true
	mBlockList["09122447108"] = true
	mBlockList["989122721634"] = true
	mBlockList["09122721634"] = true
	mBlockList["989122794330"] = true
	mBlockList["09122794330"] = true
	mBlockList["989122812811"] = true
	mBlockList["09122812811"] = true
	mBlockList["989122822823"] = true
	mBlockList["09122822823"] = true
	mBlockList["989122874408"] = true
	mBlockList["09122874408"] = true
	mBlockList["989122940601"] = true
	mBlockList["09122940601"] = true
	mBlockList["989123483812"] = true
	mBlockList["09123483812"] = true
	mBlockList["989123546689"] = true
	mBlockList["09123546689"] = true
	mBlockList["989123664066"] = true
	mBlockList["09123664066"] = true
	mBlockList["989123871097"] = true
	mBlockList["09123871097"] = true
	mBlockList["989124040640"] = true
	mBlockList["09124040640"] = true
	mBlockList["989124167569"] = true
	mBlockList["09124167569"] = true
	mBlockList["989124167569"] = true
	mBlockList["09124167569"] = true
	mBlockList["989124441386"] = true
	mBlockList["09124441386"] = true
	mBlockList["989124443271"] = true
	mBlockList["09124443271"] = true
	mBlockList["989124596073"] = true
	mBlockList["09124596073"] = true
	mBlockList["989125063543"] = true
	mBlockList["09125063543"] = true
	mBlockList["989125089676"] = true
	mBlockList["09125089676"] = true
	mBlockList["989125206642"] = true
	mBlockList["09125206642"] = true
	mBlockList["989125349636"] = true
	mBlockList["09125349636"] = true
	mBlockList["989126900023"] = true
	mBlockList["09126900023"] = true
	mBlockList["989127755816"] = true
	mBlockList["09127755816"] = true
	mBlockList["989128018548"] = true
	mBlockList["09128018548"] = true
	mBlockList["989128335503"] = true
	mBlockList["09128335503"] = true
	mBlockList["989128475975"] = true
	mBlockList["09128475975"] = true
	mBlockList["989129342383"] = true
	mBlockList["09129342383"] = true
	mBlockList["989131118028"] = true
	mBlockList["09131118028"] = true
	mBlockList["989131183892"] = true
	mBlockList["09131183892"] = true
	mBlockList["989132176815"] = true
	mBlockList["09132176815"] = true
	mBlockList["989132518595"] = true
	mBlockList["09132518595"] = true
	mBlockList["989133039745"] = true
	mBlockList["09133039745"] = true
	mBlockList["989133138367"] = true
	mBlockList["09133138367"] = true
	mBlockList["989153110979"] = true
	mBlockList["09153110979"] = true
	mBlockList["989179928657"] = true
	mBlockList["09179928657"] = true
	mBlockList["989191003321"] = true
	mBlockList["09191003321"] = true
	mBlockList["989192381823"] = true
	mBlockList["09192381823"] = true
	mBlockList["989195544557"] = true
	mBlockList["09195544557"] = true
	mBlockList["989382968154"] = true
	mBlockList["09382968154"] = true
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

func SetPattern(p string) {
	Pattern = p
}

func RelayRoutes(router *gin.RouterGroup) {
	InitializeBlockList()
	{
		router.POST("/send", SendSMS)
		router.GET("/credit", GetCredit)
		//router.GET("/status", SendStatus)
	}
}
func MakeRequest(mobile string, code string) string {
	request := fmt.Sprintf("http://robots.rahco.ir/api/proxy/send?username=khadamati1400&password=WbUqSBo&receiver_number=%s&pattern_id=%s&pattern_params[]=%s&token=rVW9HmLH41RjA5PywpuHGdfXODzbQo", mobile, Pattern, code)
	//request := fmt.Sprintf("http://localhost:3000/send")
	return request
}

func DoneAsync(mobile string, code string, generatedMsgId string) chan int {
	r := make(chan int)
	go func() {
		request := MakeRequest(mobile, code)

		//fmt.Println("request : ", request)
		resp, err1 := client.Get(request)
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

func IsBlocked(mobile string) bool {
	_, check := mBlockList[mobile]
	return check
}
func SendSMS(c *gin.Context) {
	//if utility.HasCredit() == false {
	//	utility.Log("WARNING", "credit error")
	//	c.JSON(http.StatusBadRequest, CreateErrorMessage("credit error"))
	//	return
	//}
	//c.JSON(http.StatusBadRequest, CreateErrorMessage("has credit"))
	//return
	mobile := c.PostForm("receiver_number")
	code := c.PostForm("code")
	if mobile == "" || code == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("some params is required!"))
		return
	}

	if _, err := strconv.Atoi(code); err != nil {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("only digit is acceptable"))
		return
	}

	if IsBlocked(mobile) {
		utility.Log("ERROR", "Blocked : ", mobile)
		c.JSON(http.StatusBadRequest, CreateErrorMessage("blocked"))
		return
	}

	if !utility.CheckRelay(mobile) {
		utility.Log("ERROR", "Blocked : ", mobile, " by history")
		c.JSON(http.StatusBadRequest, CreateErrorMessage("history"))
		return
	}
	//utility.DecreaseCredit()
	utility.IncreaseCredit()

	rand.Seed(time.Now().UnixNano())
	randomNum := random(10000000, 90000000)
	generatedMessageId := fmt.Sprintf("%d", randomNum)
	c.JSON(http.StatusOK, gin.H{
		"data": generatedMessageId,
	})
	DoneAsync(mobile, code, generatedMessageId)

}

func GetCredit(c *gin.Context) {
	dat, _ := ioutil.ReadFile("credit.txt")

	c.JSON(http.StatusOK, gin.H{
		"credit": string(dat),
	})

}
