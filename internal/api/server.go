package api

import (
	"fmt"
	"jenkins-pipeline-manager/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer(port int, pipelineService *service.PipelineService) {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	r.POST("/pipelines", func(c *gin.Context) {
		var req struct {
			Name            string `json:"name"`
			GitRepo         string `json:"git_repo"`
			Branch          string `json:"branch"`
			JenkinsfilePath string `json:"jenkinsfile_path"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Received request to create pipeline:", req)

		err := pipelineService.CreatePipeline(req.Name, req.GitRepo, req.Branch, req.JenkinsfilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Pipeline created successfully vinay"})
	})

	r.Run(fmt.Sprintf(":%d", port))
}
