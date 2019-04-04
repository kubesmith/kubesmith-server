package signup

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/database/models"
	"github.com/kubesmith/kubesmith-server/src/server"
)

type PostData struct {
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
	Type      string `json:"json" form:"type"`
}

type PostError struct {
	code    int
	message string
}

func (e *PostError) Error() string {
	return e.message
}

func (e *PostError) Code() int {
	return e.code
}

type SignupPostHandler struct {
	Data PostData
}

func (h *SignupPostHandler) Validate() *PostError {
	return &PostError{
		code:    http.StatusBadRequest,
		message: "bad request",
	}
}

func (h *SignupPostHandler) Process() (*models.User, *PostError) {
	if err := h.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
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
		spew.Dump(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, user)
}
