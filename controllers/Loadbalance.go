package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	_ "os/exec"
	"regexp"
	_ "runtime"
	"strings"
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

func CreateSuccessMessage(message string) gin.H {
	ret := gin.H{
		"success": true,
		"message": message,
	}
	return ret
}

func CreateErrorMessage(message string) gin.H {
	ret := gin.H{
		"success": false,
		"message": message,
	}
	return ret
}

func Convert2Instance(income string) []Instance {
	income = strings.Replace(income, "\nid", "#id", -1)
	income = strings.Replace(income, "\n\t", ",", -1)
	income = strings.Replace(income, " ", "", -1)
	income = strings.Replace(income, "\n", "", -1)

	splitData := strings.Split(income, "#")
	output := []Instance{}
	for _, data := range splitData {
		ss := strings.Split(data, ",")
		url := strings.Split(ss[1], "https")[1]
		url = "https" + url
		m := Instance{
			strings.Split(ss[0], ":")[1],
			url,
			strings.Split(ss[2], ":")[1],
			ss[3] == "enabled",
			strings.Split(ss[4], ":")[1],
			strings.Split(ss[5], ":")[1],
			ss[6] == "online"}

		output = append(output, m)

	}

	return (output)
}

func Convert2MonitoringJson(income string) gin.H {
	space := regexp.MustCompile(`\s+`)
	income = space.ReplaceAllString(income, " ")
	income = income[strings.Index(income, " ING ")+4:]
	income = strings.Replace(income, "\\n r", "", -1)
	income = strings.Replace(income, " r", "#", -1)
	splitData := strings.Split(income, "#")
	//HOSTNAME      STATE    STATUS  MEETINGS  USERS  LARGEST MEETING  VIDEOS
	var response []map[string]interface{}
	for _, data := range splitData {
		data := strings.TrimSpace(data)
		row := strings.Split(data, " ")
		if len(row) < 4 {
			continue
		}
		result := map[string]interface{}{
			"host":           row[0],
			"state":          row[1],
			"status":         row[2],
			"meetings":       row[3],
			"users":          row[4],
			"largestMeeting": row[5],
			"videos":         row[6],
		}
		response = append(response, result)
	}
	ret := gin.H{
		"success": true,
		"data":    response,
	}
	return ret
}

func Instance2Json(income []Instance) map[string]interface{} {
	var response []map[string]interface{}
	for _, data := range income {
		result := map[string]interface{}{
			"id":             data.Id,
			"url":            data.URL,
			"secret":         data.Secret,
			"enable":         data.Enable,
			"load":           data.Load,
			"loadMultiplier": data.LoadMultiplier,
			"online":         data.Online,
		}
		response = append(response, result)
	}
	ret := gin.H{
		"success": true,
		"data":    response,
	}
	return ret
}
func LoadBalanceRoutes(router *gin.RouterGroup) {
	{
		router.GET("/servers", ListServer)
		router.GET("/poll", PollServer)
		router.GET("/monitoring", Monitoring)
		router.POST("/disable", DisableServer)
		router.POST("/enable", EnableServer)
		router.POST("/panic", PanicServer)
		router.POST("/add", AddServer)
		router.POST("/delete", DeleteServer)
	}
}

func GetListServer() map[string]interface{} {
	cmd := "docker exec -i scalelite-api bundle exec rake servers"
	data, _ := exec.Command("bash", "-c", cmd).Output()
	m := Convert2Instance(string(data))
	j := Instance2Json(m)
	return j
}
func ListServer(c *gin.Context) {
	c.JSON(http.StatusOK, GetListServer())
}

func PollServer(c *gin.Context) {
	cmd := "docker exec -i scalelite-api bundle exec rake poll:all"
	exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, GetListServer())
}

func Monitoring(c *gin.Context) {
	cmd := "docker exec -i scalelite-api bundle exec rake status"
	data, _ := exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, Convert2MonitoringJson(string(data)))
}

func DisableServer(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("param(id) is required!"))
	}
	cmd := fmt.Sprintf("docker exec -i scalelite-api bundle exec rake disable[%s]", id)
	exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, GetListServer())
}

func EnableServer(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("param(id) is required!"))
	}
	cmd := fmt.Sprintf("docker exec -i scalelite-api bundle exec rake enable[%s]", id)
	exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, GetListServer())
}

func PanicServer(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("param(id) is required!"))
	}
	cmd := fmt.Sprintf("docker exec -i scalelite-api bundle exec rake panic[%s]", id)
	exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, GetListServer())
}

func DeleteServer(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("param(id) is required!"))
	}
	fmt.Println("id :", id)
	cmd := fmt.Sprintf("docker exec -i scalelite-api bundle exec rake servers:remove[%s]", id)
	exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, GetListServer())
}
func AddServer(c *gin.Context) {

	url := c.PostForm("url")
	secret := c.PostForm("secret")
	loadMultiplier := c.PostForm("loadMultiplier")
	if url == "" || secret == "" || loadMultiplier == "" {
		c.JSON(http.StatusBadRequest, CreateErrorMessage("param(id) is required!"))
	}
	cmd := fmt.Sprintf("docker exec -i scalelite-api bundle exec rake servers:add[%s,%s,%s]", url, secret, loadMultiplier)
	exec.Command("bash", "-c", cmd).Output()

	c.JSON(http.StatusOK, GetListServer())
}
