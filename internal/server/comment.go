package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/vasjaj/todo/internal/db"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// @Summary Get task comments.
// @Description Get comments by task.
// @Tags comments
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Response 200 {array} server.getCommentResponse
// @Router /tasks/{task_id}/comments [get]
func (s *Server) getTaskComments(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	comments, err := s.Database.GetComments(taskID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	res := make([]*getCommentResponse, len(comments))
	for i, comment := range comments {
		res[i] = mapCommentToResponse(&comment)
	}

	return c.JSON(http.StatusOK, res)
}

type createCommentRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// @Summary Create task comment.
// @Description Create one comment.
// @Tags comments
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Param comment body server.createCommentRequest true "Comment"
// @Router /tasks/{task_id}/comments [post]
func (s *Server) createTaskComment(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	req := &createCommentRequest{}
	if err := c.Bind(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := s.Database.CreateComment(taskID, userID, req.Title, req.Description); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

type getCommentResponse struct {
	ID          uint      `json:"comment_id"`
	TaskID      int       `json:"task_id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func mapCommentToResponse(comment *db.Comment) *getCommentResponse {
	return &getCommentResponse{
		ID:          comment.ID,
		TaskID:      comment.TaskID,
		UserID:      comment.UserID,
		Title:       comment.Title,
		Description: comment.Description,
		CreatedAt:   comment.CreatedAt,
	}
}

// @Summary Get task comment.
// @Description Get one comment.
// @Tags comments
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Param comment_id path string true "Comment ID"
// @Response 200 {object} server.getCommentResponse
// @Router /tasks/{task_id}/comments/{comment_id} [get]
func (s *Server) getTaskComment(c echo.Context) error {
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	comment, err := s.Database.GetComment(commentID)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, mapCommentToResponse(comment))
}

// @Summary Update comment.
// @Description Update one comment.
// @Tags comments
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Param comment_id path string true "Comment ID"
// @Param comment body server.createCommentRequest true "Comment"
// @Router /tasks/{task_id}/comments/{comment_id} [put]
func (s *Server) updateTaskComment(c echo.Context) error {
	comment, err := s.findComment(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	req := &createCommentRequest{}
	if err := c.Bind(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := s.Database.UpdateComment(int(comment.ID), req.Title, req.Description); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

// @Summary Delete task comment.
// @Description Delete one comment.
// @Tags comments
// @Param Authorization header string true "Authorization"
// @Param task_id path string true "Task ID"
// @Param comment_id path string true "Comment ID"
// @Router /tasks/{task_id}/comments/{comment_id} [delete]
func (s *Server) deleteTaskComment(c echo.Context) error {
	comment, err := s.findComment(c)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err := s.Database.DeleteComment(int(comment.ID)); err != nil {
		log.Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

func (s *Server) findComment(c echo.Context) (*db.Comment, error) {
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		log.Error(err)

		return nil, err
	}

	comment, err := s.Database.GetComment(commentID)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	if userID != comment.UserID {
		return nil, err
	}

	return comment, nil
}
