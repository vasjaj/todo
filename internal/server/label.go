package server

import (
	"errors"
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

	labels, err := s.db.GetLabels(userID)
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
	ID        uint   `json:"label_id"`
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

	if err := s.db.CreateLabel(userID, label.Title); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Get label.
// @Description Get label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param label_id path string true "Label ID"
// @Response 200 {object} server.getLabelResponse
// @Router /labels/{label_id} [get]
func (s *Server) getLabel(c echo.Context) error {
	label, err := s.findLabel(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, mapLabelToResponse(label))
}

// @Summary Update label.
// @Description Update label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param label_id path string true "Label ID"
// @Param label body server.createLabelRequest true "Label"
// @Response 200 {object} server.getLabelResponse
// @Router /labels/{label_id} [put]
func (s *Server) updateLabel(c echo.Context) error {
	label, err := s.findLabel(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	var req createLabelRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := s.db.UpdateLabel(int(label.ID), req.Title); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Delete label.
// @Description Delete label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param label_id path string true "Label ID"
// @Route /labels/{label_id} [delete]
func (s *Server) deleteLabel(c echo.Context) error {
	label, err := s.findLabel(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.db.DeleteLabel(int(label.ID)); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Get tasks for label.
// @Description Get tasks for label.
// @Tags labels
// @Param Authorization header string true "Authorization"
// @Param label_id path string true "Label ID"
// @Response 200 {array} server.getTaskResponse
// @Router /labels/{label_id}/tasks [get]
func (s *Server) getTasksForLabel(c echo.Context) error {
	label, err := s.findLabel(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	tasks, err := s.db.GetTasksByLabel(int(label.ID))
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

func (s *Server) findLabel(c echo.Context) (*db.Label, error) {
	labelID, err := strconv.Atoi(c.Param("label_id"))
	if err != nil {
		return nil, err
	}

	label, err := s.db.GetLabel(labelID)
	if err != nil {
		return nil, err
	}

	userID, err := getUserID(c)
	if err != nil {
		return nil, err
	}

	if label.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return label, nil
}
