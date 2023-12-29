package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/douyinDownload_go/client"
	"net/http"
)

func StartServer() {
	r := gin.Default()
	r.GET("/download", downloadVideo)
	r.Run("0.0.0.0:8080")
}

func downloadVideo(c *gin.Context) {
	url := c.Query("url")
	fileName := c.Query("fileName")
	body, err, length := client.GetBody(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	if body != nil {
		defer body.Close()
	}
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="` + fileName + `"`,
	}
	c.DataFromReader(http.StatusOK, length, "application/octet-stream",
		body, extraHeaders)
}
