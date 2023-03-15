package main

import (
	"fmt"
	"net/http"
	model2 "skywalkingdemo/pkg/model"
	"skywalkingdemo/pkg/tracerhelper"
	"skywalkingdemo/pkg/tracerhelper/ginagent"
	"skywalkingdemo/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	if tracerhelper.StartTracer(utils.GetEnv("SW_OAP_SERVER", "192.168.47.150:11800"), "test-demo3") != nil {
		fmt.Println("create gosky reporter failed!")
	}

	model2.Setup()
	defer model2.CloseAllDb()

	r := gin.New()
	r.Use(ginagent.Middleware())
	r.GET("/test", test)
	_ = r.Run(":7003")
}

func test(c *gin.Context) {
	model2.Read5ScoreLogModel{}.GetId(1, 2)
	model2.Read5WhiteListModel{}.GetId(1, 2)

	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = ""
	result["data"] = "test"
	c.JSON(http.StatusOK, result)
}
