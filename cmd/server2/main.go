package main

import (
	"fmt"
	"net/http"
	model2 "skywalkingdemo/pkg/model"
	tracerhelper2 "skywalkingdemo/pkg/tracerhelper"
	"skywalkingdemo/pkg/tracerhelper/ginagent"

	"github.com/gin-gonic/gin"
)

func main() {
	if tracerhelper2.StartTracer("192.168.47.150:11800", "test-demo2") != nil {
		fmt.Println("create gosky reporter failed!")
	}

	model2.Setup()
	defer model2.CloseAllDb()

	r := gin.New()
	r.Use(ginagent.Middleware())
	r.GET("/test", test)
	_ = r.Run(":7002")
}

func test(c *gin.Context) {
	tracerhelper2.Get("http://127.0.0.1:7003/test?name=hahaha")

	model2.Read5ScoreLogModel{}.GetId(1, 2)
	model2.Read5WhiteListModel{}.GetId(1, 2)

	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = ""
	result["data"] = "test"
	c.JSON(http.StatusOK, result)
}
