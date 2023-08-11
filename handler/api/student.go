package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI interface {
	AddStudent(c *gin.Context)
	UpdateStudent(c *gin.Context)
	DeleteStudent(c *gin.Context)
	GetStudentByID(c *gin.Context)
	GetStudentList(c *gin.Context)
	GetStudentListByClass(c *gin.Context)
}

type studentAPI struct {
	studentService service.StudentService
}

func NewStudentAPI(studentRepo service.StudentService) *studentAPI {
	return &studentAPI{studentRepo}
}

func (t *studentAPI) AddStudent(c *gin.Context) {
	var newStudent model.Mahasiswa
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.studentService.Store(&newStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add student success"})
}

func (t *studentAPI) UpdateStudent(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
	studentID := c.Param("id")

	// cek jika id kosong
	if studentID == ""{
		c.JSON(400, model.ErrorResponse{Error: "invalid student ID"})
		return
	}

	// konvert id string ke type int
	id, err := strconv.Atoi(studentID)
	if err != nil{
		c.JSON(400, model.ErrorResponse{Error: "invalid student ID"})
		return
	}

	// binding JSON dari body request ke variabel taskData
	var studentData model.Mahasiswa

	if err := c.ShouldBindJSON(&studentData); err != nil{
		c.JSON(400, model.ErrorResponse{Error: err.Error()})
		return
	}
	// mengatur id tugas yang telah diperoleh
	studentData.ID = id
	// memanggil fungsi update dari taskService untuk mengupdate
	err = t.studentService.Update(id, &studentData)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message" : "update student success", "task" : studentData,
	})
}

func (t *studentAPI) DeleteStudent(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
studentID := c.Param("id")

// cek jika id kosong
if studentID == ""{
	c.JSON(400, model.ErrorResponse{Error: "Invalid student ID"})
	return
}

// konvert id string ke type int
id, err := strconv.Atoi(studentID)
if err != nil{
	c.JSON(400, model.ErrorResponse{Error: "Invalid student ID"})
	return
}

err  = t.studentService.Delete(id)
if err != nil{
	c.JSON(500, model.ErrorResponse{Error: err.Error()})
	return
}
c.JSON(200, model.SuccessResponse{Message: "delete student success"})
}

func (t *studentAPI) GetStudentByID(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	student, err := t.studentService.GetByID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (t *studentAPI) GetStudentList(c *gin.Context) {
	result, err := t.studentService.GetList()
	if err != nil {
		c.JSON(500, model.ErrorResponse{Error: "invalid student"})
		return
	}
	c.JSON(200, result)
}

func (t *studentAPI) GetStudentListByClass(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid student ID"})
		return
	}
	result, err := t.studentService.GetStudentClass(studentID)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, result)
}
