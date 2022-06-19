package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"

	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	LoginRequest
}

// @Summary System register.
// @Description Register with username and password.
// @Tags user
// @Param register body server.RegisterRequest true "Register request"
// @Router /register [post]
func (s *Server) register(c echo.Context) error {
	req := &LoginRequest{}
	if err := c.Bind(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := validator.New().Struct(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	user, err := s.Database.CreateUser(req.Login, req.Password)
	if err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, echo.Map{
		"login": user.Login,
	})
}

type LoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"min=8"`
}

type jwtCustomClaims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

// @Summary System login.
// @Description Login with username and password.
// @Tags user
// @Param login body server.LoginRequest true "Login request"
// @Router /login [post]
func (s *Server) login(c echo.Context) error {
	req := &LoginRequest{}
	if err := c.Bind(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	if err := validator.New().Struct(req); err != nil {
		log.Error(err)

		return echo.ErrBadRequest
	}

	user, err := s.Database.GetUser(req.Login)
	if err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		log.Error(err)

		return echo.ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&jwtCustomClaims{
			user.Login,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second * time.Duration(s.jwtConfig.SecondsTTL)).Unix(),
				IssuedAt:  time.Now().Unix(),
				Subject:   strconv.Itoa(int(user.ID)),
			},
		})

	t, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func getUserID(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)

	return strconv.Atoi(claims.Subject)
}
