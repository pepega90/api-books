package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("img")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gagal upload image",
		})
		return
	}

	// save and upload
	c.SaveUploadedFile(file, "uploads/"+file.Filename)

	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:8080/v1/upload/" + file.Filename,
	})

}
