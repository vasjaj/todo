package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/vasjaj/todo/internal/db"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

// @Summary Get tasks.
// @Description Get all tasks.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Response 200 {array} server.getTaskResponse
// @Router /tasks [get]
func (s *Server) getTasks(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	tasks, err := s.Database.GetTasks(userID)
	if err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	res := make([]*getTaskResponse, len(tasks))
	for i, task := range tasks {
		res[i] = mapTaskToResponse(&task)
	}

	return c.JSON(http.StatusOK, res)
}

type createTaskRequest struct {
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	DueDate     *time.Time `json:"dueDate"`
}

// @Summary Get completed tasks.
// @Description Get all completed tasks.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Response 200 {array} server.getTaskResponse
// @Router /tasks/completed [get]
func (s *Server) getCompletedTasks(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	tasks, err := s.Database.GetCompletedTasks(userID)
	if err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	res := make([]*getTaskResponse, len(tasks))
	for i, task := range tasks {
		res[i] = mapTaskToResponse(&task)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Get uncompleted tasks.
// @Description Get all uncompleted tasks.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Response 200 {array} server.getTaskResponse
// @Router /tasks/uncompleted [get]
func (s *Server) getUncompletedTasks(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	tasks, err := s.Database.GetUncompletedTasks(userID)
	if err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	res := make([]*getTaskResponse, len(tasks))
	for i, task := range tasks {
		res[i] = mapTaskToResponse(&task)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Create task.
// @Description Create one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param task body server.createTaskRequest true "Task"
// @Router /tasks [post]
func (s *Server) createTask(c echo.Context) error {
	req := &createTaskRequest{}
	if err := c.Bind(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.Database.CreateTask(userID, req.Title, req.Description, req.DueDate); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

type getTaskResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func mapTaskToResponse(task *db.Task) *getTaskResponse {
	return &getTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
	}
}

// @Summary Get task.
// @Description Get one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task ID"
// @Response 200 {object} server.getTaskResponse
// @Router /tasks/{id} [get]
func (s *Server) getTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	task, err := s.Database.GetTask(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if task.UserID != userID {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, mapTaskToResponse(task))
}

// @Summary Update task.
// @Description Update one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task ID"
// @Param task body server.createTaskRequest true "Task"
// @Router /tasks/{id} [put]
func (s *Server) updateTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	task, err := s.Database.GetTask(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if task.UserID != userID {
		return echo.ErrUnauthorized
	}

	req := &createTaskRequest{}
	if err := c.Bind(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := s.Database.UpdateTask(taskID, req.Title, req.Description, req.DueDate); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Delete task.
// @Description Delete one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task ID"
// @Router /tasks/{id} [delete]
func (s *Server) deleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	task, err := s.Database.GetTask(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if task.UserID != userID {
		return echo.ErrUnauthorized
	}

	if err := s.Database.DeleteTask(taskID); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Complete task.
// @Description Complete one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task ID"
// @Router /tasks/{id}/complete [post]
func (s *Server) completeTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	task, err := s.Database.GetTask(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)
	}

	if task.UserID != userID {
		return echo.ErrUnauthorized
	}

	if err := s.Database.CompleteTask(taskID); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary uncomplete task.
// @Description uncomplete one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task ID"
// @Router /tasks/{id}/uncomplete [post]
func (s *Server) uncompleteTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	task, err := s.Database.GetTask(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)
	}

	if task.UserID != userID {
		return echo.ErrUnauthorized
	}

	if err := s.Database.UncompleteTask(taskID); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
