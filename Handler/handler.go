package handlers

import (
	"net/http"
	services "server-bridge/service"

	"github.com/gin-gonic/gin"
)

// UploadHandler handles file upload via SSH
func UploadHandler(c *gin.Context) {
	server := c.PostForm("server")
	user := c.PostForm("user")
	keyPath := c.PostForm("keyPath")
	localPath := c.PostForm("localPath")
	remotePath := c.PostForm("remotePath")

	client, err := services.NewSSHClient(user, keyPath, server)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := client.UploadFile(remotePath, localPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "file uploaded successfully"})
}

// DownloadHandler handles file download via SSH
func DownloadHandler(c *gin.Context) {
	server := c.Query("server")
	user := c.Query("user")
	keyPath := c.Query("keyPath")
	remotePath := c.Query("remotePath")
	localPath := c.Query("localPath")

	client, err := services.NewSSHClient(user, keyPath, server)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := client.DownloadFile(remotePath, localPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "file downloaded successfully"})
}

// ListDirectoryHandler handles listing the root directory details via SSH
func ListDirectoryHandler(c *gin.Context) {
	server := c.Query("server")
	user := c.Query("user")
	keyPath := c.Query("keyPath")

	client, err := services.NewSSHClient(user, keyPath, server)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output, err := client.ListDirectory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"directory_details": output})
}
