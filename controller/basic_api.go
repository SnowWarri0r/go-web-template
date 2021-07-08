package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-template/util"
	"net/http"
)

func ok(c *gin.Context,result util.Result) {
	c.AsciiJSON(http.StatusOK,result)
}

func success(c *gin.Context) {
	c.AsciiJSON(http.StatusOK,util.Success())
}

func successInfo(c *gin.Context,message string) {
	c.AsciiJSON(http.StatusOK,util.SuccessInfo(message))
}

func successInfoAndData(c *gin.Context,message string,data interface{}) {
	c.AsciiJSON(http.StatusOK,util.SuccessInfoAndData(message,data))
}

func successData(c *gin.Context,data interface{}) {
	c.AsciiJSON(http.StatusOK,util.SuccessData(data))
}

func fail(c *gin.Context) {
	c.AsciiJSON(http.StatusOK,util.Failure())
}

func failInfo(c *gin.Context,message string) {
	c.AsciiJSON(http.StatusOK,util.FailureInfo(message))
}

func failInfoAndData(c *gin.Context,message string,data interface{}) {
	c.AsciiJSON(http.StatusOK,util.FailureInfoAndData(message,data))
}

func failData(c *gin.Context,data interface{}) {
	c.AsciiJSON(http.StatusOK,util.FailureData(data))
}

func err(c *gin.Context) {
	c.AsciiJSON(http.StatusOK,util.Error())
}

func unauthorized(c *gin.Context) {
	c.AsciiJSON(http.StatusUnauthorized,gin.H{
		"status":0,
		"message":"Unauthorized",
	})
}
