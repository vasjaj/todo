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

	tasks, err := s.db.GetTasks(userID)
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

	tasks, err := s.db.GetCompletedTasks(userID)
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

	tasks, err := s.db.GetUncompletedTasks(userID)
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
// @Description Create one task, time format - 2018-12-10T13:49:51.141Z.
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

	if err := s.db.CreateTask(userID, req.Title, req.Description, req.DueDate); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

type getTaskResponse struct {
	ID          uint       `json:"task_id"`
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
// @Param task_id path string true "Task ID"
// @Response 200 {object} server.getTaskResponse
// @Router /tasks/{task_id} [get]
func (s *Server) getTask(c echo.Context) error {
	task, err := s.findTask(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, mapTaskToResponse(task))
}

// @Summary Update task.
// @Description Update one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Param task body server.createTaskRequest true "Task"
// @Router /tasks/{task_id} [put]
func (s *Server) updateTask(c echo.Context) error {
	task, err := s.findTask(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	req := &createTaskRequest{}
	if err := c.Bind(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := s.db.UpdateTask(int(task.ID), req.Title, req.Description, req.DueDate); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Delete task.
// @Description Delete one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Router /tasks/{task_id} [delete]
func (s *Server) deleteTask(c echo.Context) error {
	task, err := s.findTask(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.db.DeleteTask(int(task.ID)); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Complete task.
// @Description Complete one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Router /tasks/{task_id}/complete [post]
func (s *Server) completeTask(c echo.Context) error {
	task, err := s.findTask(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.db.CompleteTask(int(task.ID)); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary uncomplete task.
// @Description uncomplete one task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Router /tasks/{task_id}/uncomplete [post]
func (s *Server) uncompleteTask(c echo.Context) error {
	task, err := s.findTask(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.db.UncompleteTask(int(task.ID)); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Add label to task.
// @Description Add label to task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Param label_id path string true "Label ID"
// @Router /tasks/{task_id}/labels/{label_id} [post]
func (s *Server) addLabelToTask(c echo.Context) error {
	label, err := s.findLabel(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	task, err := s.findTask(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.db.AddLabelToTask(int(label.ID), int(task.ID)); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Remove label from task.
// @Description Remove label from task.
// @Tags tasks
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Param label_id path string true "Label ID"
// @Router /tasks/{task_id}/labels/{label_id}/ [delete]
func (s *Server) removeLabelFromTask(c echo.Context) error {
	label, err := s.findLabel(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	task, err := s.findTask(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.db.RemoveLabelFromTask(int(label.ID), int(task.ID)); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

func (s *Server) findTask(c echo.Context) (*db.Task, error) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return nil, err
	}

	task, err := s.db.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	userID, err := getUserID(c)
	if err != nil {
		return nil, err
	}

	if task.UserID != userID {
		return nil, err
	}

	return task, nil
}
