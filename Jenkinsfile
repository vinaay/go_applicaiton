@Library('my-shared-library') _

// Example usage of the shared library in an application's Jenkinsfile
buildPipeline(
  appType: 'go',
  imageName: 'example-app',
  // imageTag omitted to use BUILD_NUMBER by default
  dockerRepo: 'myregistry.example.com/myrepo',
  dockerCredsId: 'docker-credentials-id'
)
