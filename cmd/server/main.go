package main

import (
	"fmt"
	"jenkins-pipeline-manager/internal/api"
	"jenkins-pipeline-manager/internal/config"
	"jenkins-pipeline-manager/internal/jenkins"
	"jenkins-pipeline-manager/internal/service"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Println("Configuration loaded successfully", cfg.Jenkins.Token)
	// Initialize Jenkins client
	jenkinsClient := jenkins.NewJenkinsClient(cfg.Jenkins.URL, cfg.Jenkins.Username, cfg.Jenkins.Token)

	fmt.Println("Jenkins client initialized successfully", jenkinsClient)
	// Initialize services
	pipelineService := service.NewPipelineService(jenkinsClient)

	fmt.Println("Pipeline service initialized successfully")

	// Start API server
	api.StartServer(cfg.Server.Port, pipelineService)
}
