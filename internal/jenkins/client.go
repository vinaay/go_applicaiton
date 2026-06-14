package jenkins

import (
	"fmt"
	"net/http"
	"strings"
)

type JenkinsClient interface {
	ValidateConnection() error
	CreatePipelineJob(name, gitRepo, branch, jenkinsfilePath string) error
}

type client struct {
	baseURL  string
	username string
	token    string
}

func NewJenkinsClient(baseURL, username, token string) JenkinsClient {
	return &client{
		baseURL:  baseURL,
		username: username,
		token:    token,
	}
}

func (c *client) ValidateConnection() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/json", c.baseURL), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.username, c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to connect to Jenkins: %s", resp.Status)
	}

	return nil
}

func (c *client) CreatePipelineJob(name, gitRepo, branch, jenkinsfilePath string) error {
	jobConfig := fmt.Sprintf(`
        <flow-definition>
            <description>Pipeline job for %s</description>
            <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition">
                <script>pipeline {
                    agent any
                    stages {
                        stage("Checkout") {
                            steps {
                                git url: "%s", branch: "%s"
                            }
                        }
                        stage("Build") {
                            steps {
                                sh 'echo Building...'
                            }
                        }
                    }
                }</script>
                <sandbox>true</sandbox>
            </definition>
        </flow-definition>`, name, gitRepo, branch)

	req, err := http.NewRequest("POST", c.baseURL+"/createItem?name="+name, strings.NewReader(jobConfig))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(c.username, c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to Jenkins server: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusForbidden {
		return fmt.Errorf("authentication failed: Jenkins server returned status: %s", resp.Status)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create pipeline: Jenkins server returned status: %s", resp.Status)
	}

	return nil
}
