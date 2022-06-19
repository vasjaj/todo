package server

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/vasjaj/todo/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vasjaj/todo/internal/db"
)

type Server struct {
	*echo.Echo
	*db.Database
	jwtConfig
	listen string
}

type jwtConfig struct {
	Secret     string `json:"secret"`
	SecondsTTL int    `json:"seconds_ttl"`
}

// @title TODO API
// @version 1.0
// @description This is a sample server TODO server.

// @host 127.0.0.1:8080
// @BasePath /
func New(database *db.Database, conf *config.Config) *Server {
	srv := &Server{echo.New(), database, jwtConfig{conf.Server.JWT.Secret, conf.Server.JWT.SecondsTTL}, conf.Server.Listen}

	srv.Use(middleware.Logger())
	srv.Use(middleware.Recover())

	srv.GET("/swagger/*", echoSwagger.WrapHandler)
	srv.POST("/register", srv.register)
	srv.POST("/login", srv.login)

	restricted := srv.Group("")
	restricted.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(srv.jwtConfig.Secret), Claims: &jwtCustomClaims{},
	}))

	tasks := restricted.Group("/tasks")
	tasks.GET("", srv.getTasks)
	tasks.GET("/comlpeted", srv.getCompletedTasks)
	tasks.GET("/uncompleted", srv.getUncompletedTasks)
	tasks.POST("", srv.createTask)
	tasks.GET("/:task_id", srv.getTask)
	tasks.PUT("/:task_id", srv.updateTask)
	tasks.DELETE("/:task_id", srv.deleteTask)
	tasks.POST("/:task_id/complete", srv.completeTask)
	tasks.POST("/:task_id/uncomplete", srv.uncompleteTask)

	taskLabels := tasks.Group("/:task_id/labels")
	taskLabels.POST("/:label_id", srv.addLabelToTask)
	taskLabels.DELETE("/:label_id", srv.removeLabelFromTask)

	comments := tasks.Group("/:task_id/comments")
	comments.GET("", srv.getTaskComments)
	comments.POST("", srv.createTaskComment)
	comments.GET("/:comment_id", srv.getTaskComment)
	comments.PUT("/:comment_id", srv.updateTaskComment)
	comments.DELETE("/:comment_id", srv.deleteTaskComment)

	labels := restricted.Group("/labels")
	labels.GET("", srv.getLabels)
	labels.GET("/tasks", srv.getTasksForLabel)
	labels.POST("", srv.createLabel)
	labels.GET("/:label_id", srv.getLabel)
	labels.PUT("/:label_id", srv.updateLabel)
	labels.DELETE("/:label_id", srv.deleteLabel)

	return srv
}

func (s *Server) Run() error {
	return s.Start(s.listen)
}
