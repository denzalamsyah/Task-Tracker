package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassAPI interface {
	AddClass(c *gin.Context)
	UpdateClass(c *gin.Context)
	DeleteClass(c *gin.Context)
	GetClassByID(c *gin.Context)
	GetClassList(c *gin.Context)
}

type classAPI struct {
	classService service.ClassService
}

func NewClassAPI(classRepo service.ClassService) *classAPI {
	return &classAPI{classRepo}
}

func (ct *classAPI) AddClass(c *gin.Context) {
	var newClass model.Class
	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.classService.Store(&newClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add class success"})
}

func (ct *classAPI) UpdateClass(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
	classID := c.Param("id")

	// cek jika id kosong
	if classID == ""{
		c.JSON(400, model.ErrorResponse{Error: "invalid Class ID"})
		return
	}

	// konvert id string ke type int
	id, err := strconv.Atoi(classID)
	if err != nil{
		c.JSON(400, model.ErrorResponse{Error: "invalid Class ID"})
		return
	}
	var classData model.Class

	if err := c.ShouldBindJSON(&classData); err != nil{
		c.JSON(400, model.ErrorResponse{Error: err.Error()})
		return
	}

	classData.ID = id

	err = ct.classService.Update(id, classData)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
	}

	c.JSON(200, gin.H{
		"message" : "class update success", "class" : classData,
	})
}

func (ct *classAPI) DeleteClass(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
	classID := c.Param("id")

	// cek jika id kosong
	if classID == ""{
		c.JSON(400, model.ErrorResponse{Error: "invalid Class ID"})
		return
	}

	// konvert id string ke type int
	id, err := strconv.Atoi(classID)
	if err != nil{
		c.JSON(400, model.ErrorResponse{Error: "invalid Class ID"})
		return
	}

	err = ct.classService.Delete(id)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: "delete failed"})
		return
	}

	c.JSON(200, gin.H{
		"message" : "class delete success",
	})
}

func (ct *classAPI) GetClassByID(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid class ID"})
		return
	}

	class, err := ct.classService.GetByID(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, class)
}

func (ct *classAPI) GetClassList(c *gin.Context) {
	result, err := ct.classService.GetList()
	if err != nil {
		c.JSON(500, model.ErrorResponse{Error: "get class failed"})
		return
	}

	c.JSON(200, result)// TODO: answer here
}
