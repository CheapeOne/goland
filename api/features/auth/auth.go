package auth

import (
	"errors"
	"net/http"

	"github.com/cheapeone/goland/api/features/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	us *users.UserStore
}

func NewAuthHandler(db *sqlx.DB) *AuthHandler {
	us := &users.UserStore{DB: db}
	return &AuthHandler{us}
}

func (h *AuthHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := h.us.FindByEmail(email)
	if err != nil || user == nil {
		return errors.New("No user")
	}

	if !ComparePassword(password, user.Password) {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	t, _ := token.SignedString([]byte("secret"))

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (h *AuthHandler) Signup(c echo.Context) error {
	return nil
}

func HashPassword(originalPassword string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(originalPassword), 10)
	return string(pass), err
}

func ComparePassword(password, hash string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil {
		return true
	}
	return false
}
