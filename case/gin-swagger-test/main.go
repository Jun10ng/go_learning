package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "gin-swagger/docs"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name hanyun
// @contact.url
// @contact.email 1355081829@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/v1", Index)
	router.Run()
}

// @Tags 首页
// @Id 1
// @Summary 获得首页数据
// @Description | 项目 | 价格 | 数量 |
// @Description | :-------- | --------:| :--: |
// @Description | iPhone | 6000 元 | 5 |
// @Description | iPad | 3800 元 | 12 |
// @Description | iMac | 10000 元 | 234 |
// @Produce  json
// @Security ApiKeyAuth
// @Param id query string true "用户id"
// @Success 200 {object} main.JsonResult "成功"
// @Failure 400 {object} main.JsonResult "失败"
// @Router /api/v1 [get]
func Index(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	rep := &JsonResult{}
	data := &CommonResponse{
		ID: id,
	}
	rep.Data = data
	c.JSON(200, rep)
}

// JsonResult 封装
type JsonResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
type CommonResponse struct {
	ID string `json:"id"`
}
