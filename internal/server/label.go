package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/vasjaj/todo/internal/db"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// @Summary Get labels.
// @Description Get all labels.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Response 200 {array} server.getLabelResponse
// @Router /labels [get]
func (s *Server) getLabels(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	labels, err := s.Database.GetLabels(userID)
	if err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	res := make([]*getLabelResponse, len(labels))
	for i, label := range labels {
		res[i] = mapLabelToResponse(&label)
	}

	return c.JSON(http.StatusOK, res)
}

type getLabelResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
}

func mapLabelToResponse(label *db.Label) *getLabelResponse {
	return &getLabelResponse{
		ID:        label.ID,
		Title:     label.Title,
		CreatedAt: label.CreatedAt.Format(time.RFC3339),
	}
}

type createLabelRequest struct {
	Title string `json:"title" validate:"required"`
}

// @Summary Create label.
// @Description Create label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param label body server.createLabelRequest true "Label"
// @Response 200 {object} server.getLabelResponse
// @Router /labels [post]
func (s *Server) createLabel(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	var label createLabelRequest
	if err := c.Bind(&label); err != nil {
		log.Error(err)
	}

	if err := s.Database.CreateLabel(userID, label.Title); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Get label.
// @Description Get label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param id path string true "Label ID"
// @Response 200 {object} server.getLabelResponse
// @Router /labels/{id} [get]
func (s *Server) getLabel(c echo.Context) error {
	labelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	label, err := s.Database.GetLabel(labelID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if label.UserID != userID {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, mapLabelToResponse(label))
}

// @Summary Update label.
// @Description Update label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param id path string true "Label ID"
// @Param label body server.createLabelRequest true "Label"
// @Response 200 {object} server.getLabelResponse
// @Router /labels/{id} [put]
func (s *Server) updateLabel(c echo.Context) error {
	labelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	label, err := s.Database.GetLabel(labelID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if label.UserID != userID {
		return echo.ErrUnauthorized
	}

	var req createLabelRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := s.Database.UpdateLabel(labelID, req.Title); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Delete label.
// @Description Delete label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param id path string true "Label ID"
// @Route /labels/{id} [delete]
func (s *Server) deleteLabel(c echo.Context) error {
	labelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	label, err := s.Database.GetLabel(labelID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if label.UserID != userID {
		return echo.ErrUnauthorized
	}

	if err := s.Database.DeleteLabel(labelID); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Get tasks for label.
// @Description Get tasks for label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param id path string true "Label ID"
// @Response 200 {array} server.getTaskResponse
// @Router /labels/{id}/tasks [get]
func (s *Server) getTasksForLabel(c echo.Context) error {
	labelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	label, err := s.Database.GetLabel(labelID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if label.UserID != userID {
		return echo.ErrUnauthorized
	}

	tasks, err := s.Database.GetTasksForLabel(labelID)
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

// @Summary Add label to task.
// @Description Add label to task.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param id path string true "Label ID"
// @Param task_id path string true "Task ID"
// @Route /labels/{id}/tasks/{task_id} [post]
func (s *Server) addLabelToTask(c echo.Context) error {
	labelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	label, err := s.Database.GetLabel(labelID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if label.UserID != userID {
		return echo.ErrUnauthorized
	}

	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	task, err := s.Database.GetTask(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if task.UserID != userID {
		return echo.ErrUnauthorized
	}

	if err := s.Database.AddLabelToTask(labelID, taskID); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Remove label from task.
// @Description Remove label from task.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param id path string true "Label ID"
// @Param task_id path string true "Task ID"
// @Route /labels/{id}/tasks/{task_id} [delete]
func (s *Server) removeLabelFromTask(c echo.Context) error {
	labelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	label, err := s.Database.GetLabel(labelID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if label.UserID != userID {
		return echo.ErrUnauthorized
	}

	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	task, err := s.Database.GetTask(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if task.UserID != userID {
		return echo.ErrUnauthorized
	}

	if err := s.Database.RemoveLabelFromTask(labelID, taskID); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
