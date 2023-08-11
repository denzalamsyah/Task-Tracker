package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DosenAPI interface {
	AddDosen(c *gin.Context)
	UpdateDosen(c *gin.Context)
	DeleteDosen(c *gin.Context)
	GetDosenByID(c *gin.Context)
	GetDosenList(c *gin.Context)
	GetDosenListByMatkul(c *gin.Context)
}

type dosenAPI struct {
	dosenService service.DosenService
}

func NewDosenAPI(dosenRepo service.DosenService) *dosenAPI {
	return &dosenAPI{dosenRepo}
}

func (t *dosenAPI) AddDosen(c *gin.Context) {
	var newDosen model.Dosen
	if err := c.ShouldBindJSON(&newDosen); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.dosenService.Store(&newDosen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add dosen success"})
}

func (t *dosenAPI) UpdateDosen(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
	dosenID := c.Param("id")

	// cek jika id kosong
	if dosenID == ""{
		c.JSON(400, model.ErrorResponse{Error: "invalid dosen ID"})
		return
	}

	// konvert id string ke type int
	id, err := strconv.Atoi(dosenID)
	if err != nil{
		c.JSON(400, model.ErrorResponse{Error: "invalid dosen ID"})
		return
	}

	// binding JSON dari body request ke variabel taskData
	var dosenData model.Dosen

	if err := c.ShouldBindJSON(&dosenData); err != nil{
		c.JSON(400, model.ErrorResponse{Error: err.Error()})
		return
	}
	// mengatur id tugas yang telah diperoleh
	dosenData.ID = id
	// memanggil fungsi update dari taskService untuk mengupdate
	err = t.dosenService.Update(id, &dosenData)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message" : "update dosen success", "dosen" : dosenData,
	})
}

func (t *dosenAPI) DeleteDosen(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
dosenID := c.Param("id")

// cek jika id kosong
if dosenID == ""{
	c.JSON(400, model.ErrorResponse{Error: "Invalid dosen ID"})
	return
}

// konvert id string ke type int
id, err := strconv.Atoi(dosenID)
if err != nil{
	c.JSON(400, model.ErrorResponse{Error: "Invalid dosen ID"})
	return
}

err  = t.dosenService.Delete(id)
if err != nil{
	c.JSON(500, model.ErrorResponse{Error: err.Error()})
	return
}
c.JSON(200, model.SuccessResponse{Message: "delete dosen success"})
}

func (t *dosenAPI) GetDosenByID(c *gin.Context) {
	dosenID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid dosen ID"})
		return
	}

	dosen, err := t.dosenService.GetByID(dosenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dosen)
}

func (t *dosenAPI) GetDosenList(c *gin.Context) {
	result, err := t.dosenService.GetList()
	if err != nil {
		c.JSON(500, model.ErrorResponse{Error: "invalid dosen"})
		return
	}
	c.JSON(200, result)
}

func (t *dosenAPI) GetDosenListByMatkul(c *gin.Context) {
	dosenID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid dosen ID"})
		return
	}
	result, err := t.dosenService.GetDosenMatkul(dosenID)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, result)
}
