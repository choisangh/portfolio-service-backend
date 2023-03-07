package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/choisangh/board-crud-backend/pkg/model"
	"github.com/choisangh/board-crud-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (apis *APIs) CreateUser(c *gin.Context) {
	req := &model.User{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	req.Password = string(hashedPassword)
	res, err := apis.db.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id",
		})
		return
	}

	res, err := apis.db.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetUserList(c *gin.Context) {
	log.Println(c.Params)
	c.Header("Access-Control-Allow-Origin", "*")
	res, err := apis.db.GetUserList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{Res: "Server error"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetUserCurrent(c *gin.Context) {
	currentUserId, _ := c.Get("currentUserId")
	fmt.Println(currentUserId)
	c.Header("Access-Control-Allow-Origin", "*")
	userId, ok := currentUserId.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id",
		})
		return
	}

	currentUserInfo, err := apis.db.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, currentUserInfo)
}

func (apis *APIs) LoginUser(c *gin.Context) {
	req := &model.User{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := apis.db.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	loginUser := gin.H{
		"token":        token,
		"id":           user.ID,
		"email":        user.Email,
		"name":         user.Name,
		"description":  user.Description,
		"errorMessage": nil,
	}

	c.JSON(http.StatusOK, loginUser)
}
