package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhftu/api-multi/service"
)

type requestData struct {
	Index string `json:"index"`
}

func GetWorkerResultHandler(c *gin.Context) {
	var reqData requestData
	if err := c.ShouldBindJSON(&reqData); err != nil {

		c.JSON(http.StatusBadRequest, errors.New("Bad request"))
		return
	}
	fib, err := service.SaveFib(reqData.Index)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	fmt.Println("index: ", reqData.Index)
	fmt.Println("value: ", fib.Value)
	c.JSON(http.StatusOK, fib)
}

func GetLatestFibHandler(c *gin.Context) {
	result, err := service.GetLatest()
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetFibFromDB(c *gin.Context) {
	var reqData requestData
	if err := c.ShouldBindJSON(&reqData); err != nil {

		c.JSON(http.StatusBadRequest, errors.New("Bad request"))
		return
	}
	result, err := service.GetFib(reqData.Index)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllFinController(c *gin.Context) {

	result, err := service.GetAllFib()
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteFibHandler(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteFib(id); err != nil {
		c.JSON(http.StatusNotImplemented, err)
	}
	c.JSON(http.StatusOK, id)
}
