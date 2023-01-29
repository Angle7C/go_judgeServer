package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testts/config"
	"testts/dal/model"
	"testts/service"
	"testts/untils"
)

func Init() {
	r := config.GetGin()
	r.POST("/run", func(context *gin.Context) {
		var run = new(model.Run)

		if err := untils.JSONDecode(context.Request.Body, run); err != nil {
			context.JSON(http.StatusOK, model.Ok("参数传递错误"))
		} else {
			records := service.Run(*run, config.GetConfig())
			context.JSON(http.StatusOK, model.OkWithList("所有执行点执行完毕", records))

		}
	})
}
