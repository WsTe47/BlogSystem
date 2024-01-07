package api

import (
	"BlogSystem/internal/app/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

var UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func UserRegisterHandler(c *gin.Context, user *logic.UserOperate) {

	if err := c.ShouldBindJSON(&UserRegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := user.Register(c, UserRegisterRequest.Username, UserRegisterRequest.Password, UserRegisterRequest.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userID": userID})
}
