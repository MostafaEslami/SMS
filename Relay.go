package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "os/exec"
	"regexp"
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
var checkRedis = "False"

var client = http.Client{
	Timeout: 10 * time.Second,
}

var mBlockList = make(map[string]bool)

func InitializeBlockList() {
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
func CheckRedis(p string) {
	checkRedis = p
}

func RelayRoutes(router *gin.RouterGroup) {
	InitializeBlockList()
	{
		router.POST("/send", SendSMS)
		router.GET("/credit", GetCreditFromFile)
		//router.GET("/status", SendStatus)
	}
}

//func MakeRequest(mobile string, code string) string {
//	//request := fmt.Sprintf("http://robots.rahco.ir/api/proxy/send?username=khadamati1400&password=WbUqSBo&receiver_number=%s&pattern_id=%s&pattern_params[]=%s&token=rVW9HmLH41RjA5PywpuHGdfXODzbQo", mobile, Pattern, code)
//	//request := fmt.Sprintf("http://localhost:3000/send")
//	request := fmt.Sprintf("https://api.bitel.rest/api/v2/sms/single")
//
//	return request
//}

func DoneAsync(mobile string, code string, generatedMsgId string) chan int {
	r := make(chan int)
	go func() {
		//request := MakeRequest(mobile, code)
		//api := kavenegar.New("32484B49457845756E635A73623950616D536F4932654D707268534C3070657A32674241706B4A6E5032513D")
		//sender := ""
		//receptor := []string{"98", "9132957573"}
		//message := "سلام و درود"
		//if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		//	switch err := err.(type) {
		//	case *kavenegar.APIError:
		//		fmt.Println(err.Error())
		//	case *kavenegar.HTTPError:
		//		fmt.Println(err.Error())
		//	default:
		//		fmt.Println(err.Error())
		//	}
		//} else {
		//	for _, r := range res {
		//		fmt.Println("MessageID 	= ", r.MessageID)
		//		fmt.Println("Status    	= ", r.Status)
		//		//...
		//	}
		//}
		//receptor := mobile
		//template := "digiyab"
		//token := code
		//params := &kavenegar.VerifyLookupParam{}
		//if res, err := api.Verify.Lookup(receptor, template, token, params); err != nil {
		//	switch err := err.(type) {
		//	case *kavenegar.APIError:
		//		fmt.Println(err.Error())
		//	case *kavenegar.HTTPError:
		//		fmt.Println(err.Error())
		//	default:
		//		fmt.Println(err.Error())
		//	}
		//
		//	cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: "FAILED"}
		//	LogCDR(cdr)
		//	Log("ERROR", cdr.Log())
		//} else {
		//	//fmt.Println("MessageID 	= ", res.MessageID)
		//	//fmt.Println("Status    	= ", res.Status)
		//	//...
		//	s := fmt.Sprintf("%d", res.MessageID)
		//	cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: s}
		//	LogCDR(cdr)
		//	Log("INFO", cdr.Log())
		//}
		//fmt.Println("request : ", request) //5sms
		//resp, err1 := client.Get(request)
		//if resp == nil || resp.Body == nil || err1 != nil {
		//	cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: "FAILED"}
		//	LogCDR(cdr)
		//	Log("ERROR", cdr.Log())
		//	return
		//}
		//defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	Log("ERROR", err)
		//	cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: "FAILEDD"}
		//	LogCDR(cdr)
		//	Log("ERROR", cdr.Log())
		//	return
		//}
		//
		//s := fmt.Sprintf("%s", body)
		//cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: s}
		//LogCDR(cdr)
		//Log("INFO", cdr.Log())

		message := fmt.Sprintf("به اپلیکیشن خوش آمدید\n کد یکبار صرف شما : %s", code)
		bb := fmt.Sprintf("{'message': '%s', 'phoneNumber': '%s'}", message, mobile)
		myJson := bytes.NewBuffer([]byte(bb))
		req, err := http.NewRequest(http.MethodPost, "https://api.bitel.rest/api/v2/sms/single", myJson)
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMDk5OCIsImV4cCI6MTk0MTQ0NTAyNywiaXNzIjoiQml0ZWwiLCJhdWQiOiJCaXRlbCJ9.DbcrTX6MP_8p8_ReJM_HusXRGosu29l3DFl1yunRCac")
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		if resp == nil || resp.Body == nil || err != nil {
			cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: "FAILED"}
			LogCDR(cdr)
			Log("ERROR", cdr.Log())
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Log("ERROR", err)
			cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: "FAILEDD"}
			LogCDR(cdr)
			Log("ERROR", cdr.Log())
			return
		}

		s := fmt.Sprintf("%s", body)
		cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMsgId, MessageId: s}
		LogCDR(cdr)
		Log("INFO", cdr.Log())

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
	rand.Seed(time.Now().UnixNano())
	randomNum := random(10000000, 90000000)

	generatedMessageId := fmt.Sprintf("%d", randomNum)
	mobile := c.PostForm("receiver_number")
	code := c.PostForm("code")
	if mobile == "" || code == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("some params is required!"))
		return
	}

	match, _ := regexp.MatchString("^(\\+98|98|0?)9(1\\d|9[0-4])[0-9]{7}$", mobile)
	if !match {
		Log("ERROR", "only mci ", mobile)
		c.JSON(http.StatusBadRequest, CreateErrorMessage("only mci"))
		return
	}
	if _, err := strconv.Atoi(code); err != nil {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("only digit is acceptable"))
		return
	}

	if IsBlocked(mobile) {
		Log("ERROR", "Blocked : ", mobile)
		c.JSON(http.StatusBadRequest, CreateErrorMessage("blocked"))
		return
	}

	if checkRedis != "False" && mobile != "09132957573" && !CheckRelay(mobile) {
		Log("ERROR", "Blocked : ", mobile, " by history")
		cdr := CDR{Number: mobile, Code: code, MyMessageId: generatedMessageId, MessageId: "HISTORY"}
		LogCDR(cdr)
		Log("WARNING", cdr.Log())
		//c.JSON(http.StatusBadRequest, CreateErrorMessage("history"))

		c.JSON(http.StatusOK, gin.H{
			"data": generatedMessageId,
		})

		return
	}
	//utility.DecreaseCredit()
	IncreaseCredit()

	c.JSON(http.StatusOK, gin.H{
		"data": generatedMessageId,
	})

	DoneAsync(mobile, code, generatedMessageId)

}

func GetCreditFromFile(c *gin.Context) {
	dat, _ := ioutil.ReadFile("credit.txt")

	c.JSON(http.StatusOK, gin.H{
		"credit": string(dat),
	})

}
