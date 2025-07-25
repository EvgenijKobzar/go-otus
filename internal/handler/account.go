package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"otus/internal/lib/hash"
	"otus/internal/lib/jwt"
	"otus/internal/model"
	"otus/internal/repository/postgres/gorm"
)

type AccountLoginRequest struct {
	Login    string `json:"login" example:"ekobzar"`
	Password string `json:"password" example:"123456"`
}
type AccountRegisterRequest struct {
	Name      string `json:"name" example:"Evgenij"`
	FirstName string `json:"first_name" example:"Kobzar"`
	LastName  string `json:"last_name" example:""`
	Login     string `json:"login" example:"ekobzar"`
	Password  string `json:"password" example:"123456"`
}
type AccountItemResponse struct {
	Result struct {
		Item model.Account
	}
}
type AccountItemsResponse struct {
	Result struct {
		Items []model.Account
	}
}

// GetAccount godoc
// @Summary Get account by ID
// @Description Get detailed information about a account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} AccountItemResponse "Successfully retrieved account"
// @Failure 400 {object} ErrorResponse "Not found"
// @Router /otus.account.get/{id} [get]
func (h *Handler[T]) GetAccount(c *gin.Context) {
	h.getAction(c)
}

// GetListAccount godoc
// @Summary Get accounts
// @Description Get list information about account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} AccountItemsResponse "Successfully retrieved account"
// @Router /otus.account.list [get]
func (h *Handler[T]) GetListAccount(c *gin.Context) {
	h.getListAction(c)
}

// RegisterAccount godoc
// @Summary Register new account
// @Description Add a new account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body AccountRegisterRequest true "Account data"
// @Router /otus.account.register [post]
func RegisterAccount(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := hash.Make(account.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to secure password"})
		return
	}

	account.Password = hashedPassword

	repo := gorm.NewRepository[*model.Account]()
	if err = repo.Save(&account); err != nil {
		c.JSON(501, gin.H{"error": "Save hash failed"})
	}

	c.JSON(201, gin.H{"status": "user created"})
}

// DeleteAccount godoc
// @Summary Delete account
// @Description Delete a account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.account.delete/{id} [delete]
func (h *Handler[T]) DeleteAccount(c *gin.Context) {
	h.deleteAction(c)
}

// LoginAccount godoc
// @Summary Get token by UserName
// @Description Get token by UserName
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body AccountLoginRequest true "Account data"
// @Router /otus.account.login [post]
func LoginAccount(c *gin.Context) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	account, err := getAccountByLogin(credentials.Login)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if !hash.Check(credentials.Password, account.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Генерируем JWT токен...
	var token string
	if token, err = jwt.MakeByLogin(account.Login); err != nil {
		c.JSON(500, gin.H{"error": "Failed to secure password"})
	}

	c.JSON(200, gin.H{
		"token":   token,
		"expires": os.Getenv("JWT_EXPIRE_HOURS") + " hours",
	})
}

func getAccountByLogin(login string) (*model.Account, error) {
	repo := gorm.NewRepository[*model.Account]()
	items, _ := repo.GetAll()
	for _, item := range items {
		if item.Login == login {
			return item, nil
		}
	}
	return nil, errors.New("Account not found")
}
