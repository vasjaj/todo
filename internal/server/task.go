package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (s *Server) getTasks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (s *Server) createTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (s *Server) getTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (s *Server) updateTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (s *Server) deleteTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
