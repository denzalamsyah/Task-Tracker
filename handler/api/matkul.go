package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatkulAPI interface {
	AddMatkul(c *gin.Context)
	UpdateMatkul(c *gin.Context)
	DeleteMatkul(c *gin.Context)
	GetMatkulByID(c *gin.Context)
	GetMatkulList(c *gin.Context)
}

type matkulAPI struct {
	matkulService service.MatkulService
}

func NewMatkulAPI(matkulRepo service.MatkulService) *matkulAPI {
	return &matkulAPI{matkulRepo}
}

func (ct *matkulAPI) AddMatkul(c *gin.Context) {
	var newMatkul model.Matkul
	if err := c.ShouldBindJSON(&newMatkul); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.matkulService.Store(&newMatkul)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add Matkul success"})
}

func (ct *matkulAPI) UpdateMatkul(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
	matkulID := c.Param("id")

	// cek jika id kosong
	if matkulID == ""{
		c.JSON(400, model.ErrorResponse{Error: "invalid Matkul ID"})
		return
	}

	// konvert id string ke type int
	id, err := strconv.Atoi(matkulID)
	if err != nil{
		c.JSON(400, model.ErrorResponse{Error: "invalid Matkul ID"})
		return
	}
	var matkulData model.Matkul

	if err := c.ShouldBindJSON(&matkulData); err != nil{
		c.JSON(400, model.ErrorResponse{Error: err.Error()})
		return
	}

	matkulData.ID = id

	err = ct.matkulService.Update(id, matkulData)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
	}

	c.JSON(200, gin.H{
		"message" : "class update success", "matkul" : matkulData,
	})
}

func (ct *matkulAPI) DeleteMatkul(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
	matkulID := c.Param("id")

	// cek jika id kosong
	if matkulID == ""{
		c.JSON(400, model.ErrorResponse{Error: "invalid Matkul ID"})
		return
	}

	// konvert id string ke type int
	id, err := strconv.Atoi(matkulID)
	if err != nil{
		c.JSON(400, model.ErrorResponse{Error: "invalid Matkul ID"})
		return
	}

	err = ct.matkulService.Delete(id)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: "delete failed"})
		return
	}

	c.JSON(200, gin.H{
		"message" : "Matkul delete success",
	})
}

func (ct *matkulAPI) GetMatkulByID(c *gin.Context) {
	matkulID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid MatkulID"})
		return
	}

	matkul, err := ct.matkulService.GetByID(matkulID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, matkul)
}

func (ct *matkulAPI) GetMatkulList(c *gin.Context) {
	result, err := ct.matkulService.GetList()
	if err != nil {
		c.JSON(500, model.ErrorResponse{Error: "get Matkul failed"})
		return
	}

	c.JSON(200, result)// TODO: answer here
}
