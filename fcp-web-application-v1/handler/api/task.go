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
	// TODO: answer here
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		// c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = taskID

	err = t.taskService.Update(taskID, &newTask)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, model.SuccessResponse{Message: "update task success"})
	c.JSON(http.StatusOK, gin.H{"message": "update task success"})

}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	// TODO: answer here
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = t.taskService.Delete(taskID)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete task success"})
	c.JSON(http.StatusOK, gin.H{"message": "delete task success"})
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
	// TODO: answer here
	task, err := t.taskService.GetList()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	// TODO: answer here
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetTaskCategory(taskID)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}
