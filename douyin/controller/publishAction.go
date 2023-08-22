package controller

import (
	"douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishActionHandler(c *gin.Context) {
	var req models.DouyinPublishActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement video upload logic here

	res := models.DouyinPublishActionResponse{
		StatusCode: 0,
	}

	c.JSON(http.StatusOK, res)
}
