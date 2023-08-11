package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
	taskID := c.Param("id")

	// cek jika id kosong
	if taskID == ""{
		c.JSON(400, model.ErrorResponse{Error: "invalid task ID"})
		return
	}

	// konvert id string ke type int
	id, err := strconv.Atoi(taskID)
	if err != nil{
		c.JSON(400, model.ErrorResponse{Error: "invalid task ID"})
		return
	}

	// binding JSON dari body request ke variabel taskData
	var taskData model.Task

	if err := c.ShouldBindJSON(&taskData); err != nil{
		c.JSON(400, model.ErrorResponse{Error: err.Error()})
		return
	}
	// mengatur id tugas yang telah diperoleh
	taskData.ID = id
	// memanggil fungsi update dari taskService untuk mengupdate
	err = t.taskService.Update(id, &taskData)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message" : "update task success", "task" : taskData,
	})
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	// Mendapatkan ID tugas dari parameter URL
taskID := c.Param("id")

// cek jika id kosong
if taskID == ""{
	c.JSON(400, model.ErrorResponse{Error: "Invalid task ID"})
	return
}

// konvert id string ke type int
id, err := strconv.Atoi(taskID)
if err != nil{
	c.JSON(400, model.ErrorResponse{Error: "Invalid task ID"})
	return
}

err  = t.taskService.Delete(id)
if err != nil{
	c.JSON(500, model.ErrorResponse{Error: err.Error()})
	return
}
c.JSON(200, model.SuccessResponse{Message: "delete task success"})
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	result, err := t.taskService.GetList()
	if err != nil {
		c.JSON(500, model.ErrorResponse{Error: "invalid task"})
		return
	}
	c.JSON(200, result)
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid task ID"})
		return
	}
	result, err := t.taskService.GetTaskCategory(taskID)
	if err != nil{
		c.JSON(500, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, result)
}
