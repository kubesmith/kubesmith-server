package signup

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kubesmith/kubesmith-server/src/database/models"
	"github.com/kubesmith/kubesmith-server/src/server"
	"github.com/kubesmith/kubesmith-server/src/services"
	"github.com/kubesmith/kubesmith-server/src/validator"
)

type PostData struct {
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
}

type PostResponse struct {
	Code    int
	Payload interface{}
}

type SignupPostHandler struct {
	Data PostData
}

func (h *SignupPostHandler) Validate() error {
	if !validator.IsValidEmail(h.Data.Email) {
		return errors.New("email is invalid")
	}

	if !validator.IsValidPassword(h.Data.Password) {
		return errors.New("password is invalid")
	}

	return nil
}

func (h *SignupPostHandler) AccountWithEmailExists() (bool, error) {
	var account models.Account

	if err := services.GetDB().Where("email like ?", fmt.Sprintf("%%%s%%", h.Data.Email)).First(&account).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (h *SignupPostHandler) CreateUser() (*models.User, error) {
	user := models.User{
		FirstName: h.Data.FirstName,
		LastName:  h.Data.LastName,
	}

	if err := services.GetDB().Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (h *SignupPostHandler) CreateAccount(user *models.User) (*models.Account, error) {
	account := models.Account{
		UserID:   user.ID,
		Email:    h.Data.Email,
		Password: h.Data.Password,
		Type:     models.AccountTypeAccount,
	}

	if err := services.GetDB().Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (h *SignupPostHandler) Process() PostResponse {
	if err := h.Validate(); err != nil {
		return PostResponse{
			Code:    http.StatusBadRequest,
			Payload: err.Error(),
		}
	}

	exists, err := h.AccountWithEmailExists()
	if err != nil {
		fmt.Printf("an error occurred while checking for account with email: %s\n", err.Error())

		return PostResponse{
			Code: http.StatusInternalServerError,
		}
	} else if exists {
		return PostResponse{
			Code: http.StatusConflict,
		}
	}

	user, err := h.CreateUser()
	if err != nil {
		fmt.Printf("could not create user: %s\n", err.Error())

		return PostResponse{
			Code: http.StatusInternalServerError,
		}
	}

	if _, err = h.CreateAccount(user); err != nil {
		fmt.Printf("could not create account: %s\n", err.Error())

		return PostResponse{
			Code: http.StatusInternalServerError,
		}
	}

	return PostResponse{
		Code: http.StatusOK,
	}
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

	response := handler.Process()
	if response.Payload != nil {
		c.JSON(response.Code, response.Payload)
	} else {
		c.Status(response.Code)
	}
}
