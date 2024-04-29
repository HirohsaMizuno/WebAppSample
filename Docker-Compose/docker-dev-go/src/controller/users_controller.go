package controller

import (
	"docker-dev-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ユーザー一覧をJSON形式で出力
func UsersIndex(c *gin.Context) {
	users, err := model.GetUsersAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
