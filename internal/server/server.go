package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/vasjaj/todo/internal/db"
)

type Server struct {
	*echo.Echo
	*db.Database
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func New(database *db.Database) *Server {
	srv := &Server{echo.New(), database}

	srv.Use(middleware.Logger())
	srv.Use(middleware.Recover())

	srv.GET("/", srv.getTasks)
	srv.GET("/tasks", srv.getTasks)
	srv.POST("/tasks", srv.createTask)
	srv.GET("/tasks/:id", srv.getTask)
	srv.PUT("/tasks/:id", srv.updateTask)
	srv.DELETE("/tasks/:id", srv.deleteTask)

	srv.GET("/swagger/*", echoSwagger.WrapHandler)

	return srv
}

func (s *Server) Run(address string) {
	s.Logger.Fatal(s.Start(address))
}

// Handler
func (s *Server) hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
