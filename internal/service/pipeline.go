package service

import "jenkins-pipeline-manager/internal/jenkins"

type PipelineService struct {
	jenkinsClient jenkins.JenkinsClient
}

func NewPipelineService(jenkinsClient jenkins.JenkinsClient) *PipelineService {
	return &PipelineService{jenkinsClient: jenkinsClient}
}

func (s *PipelineService) CreatePipeline(name, gitRepo, branch, jenkinsfilePath string) error {
	return s.jenkinsClient.CreatePipelineJob(name, gitRepo, branch, jenkinsfilePath)
}
