package signup

import (
	"errors"
	"net/http"
	"regexp"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/database/models"
	"github.com/kubesmith/kubesmith-server/src/server"
)

type PostData struct {
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
}

type SignupPostHandler struct {
	Data PostData
}

func (h *SignupPostHandler) IsValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return len(email) <= 254 && emailRegex.MatchString(email)
}

func (h *SignupPostHandler) Validate() error {
	if !h.IsValidEmail(h.Data.Email) {
		return errors.New("email is invalid")
	}

	if utf8.RuneCountInString(h.Data.Password) < 5 {
		return errors.New("password is invalid")
	}

	return nil
}

func (h *SignupPostHandler) AccountWithEmailExists() (bool, error) {
	return false, nil
}

func (h *SignupPostHandler) CreateUser() (*models.User, error) {
	return nil, nil
}

func (h *SignupPostHandler) CreateAccount(user *models.User) (*models.Account, error) {
	return nil, nil
}

func (h *SignupPostHandler) Process() (*models.User, error) {
	if err := h.Validate(); err != nil {
		return nil, err
	}

	exists, err := h.AccountWithEmailExists()
	if err != nil {
		return nil, err
	} else if exists {
		return nil, errors.New("already registered")
	}

	user, err := h.CreateUser()
	if err != nil {
		return nil, err
	}

	if _, err = h.CreateAccount(user); err != nil {
		return nil, err
	}

	return user, nil
}

func SignupPost(server *server.Server, c *gin.Context) {
	var userData PostData

	err := c.Bind(&userData)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	handler := SignupPostHandler{
		Data: userData,
	}

	user, err := handler.Process()
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, user)
}
