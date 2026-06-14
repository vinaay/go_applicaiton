@Library('my-shared-library') _

// Example usage of the shared library in an application's Jenkinsfile
buildPipeline(
  appType: 'go',
  imageName: 'example-app',
  IMAGE_TAG : '1',
  // imageTag omitted to use BUILD_NUMBER by default
  dockerRepo: 'docker.io/vinaay/go',
  dockerCredsId: 'docker-credentials-id'
)
