pipeline {
  environment {
    IMAGE_NAME = 'payment-service'
    GITHUB_CREDENTIALS = credentials('github-credential')
  }

  stages {
    stage('Check Commit Message') {
      steps {
        script {
          def commitMessage = sh(
            script: "git log -1 --pretty=%B",
            returnStdout: true
          ).trim()

          echo "Commit Message: ${commitMessage}"
          if (commitMessage.contains("[skip ci]")) {
            echo "Skipping pipeline due to [skip ci] tag in commit message."
            currentBuild.result = 'ABORTED'
            currentBuild.delete()
            return
          }

          echo "Pipeline will continue. No [skip ci] tag found in commit message."
        }
      }
    }

    stage('Set Target Branch') {
      steps {
        script {
          echo "GIT_BRANCH: ${env.GIT_BRANCH}"
          if (env.GIT_BRANCH == 'origin/main') {
            env.TARGET_BRANCH = 'master'
          } else if (env.GIT_BRANCH == 'origin/development') {
            env.TARGET_BRANCH = 'development'
          }
        }
      }
    }

    stage('Checkout Code') {
      steps {
        script {
          def repoUrl = 'https://github.com/bwa-project-example/payment-service.git'

          checkout([$class: 'GitSCM',
            branches: [
              [name: "*/${env.TARGET_BRANCH}"]
            ],
            userRemoteConfigs: [
              [url: repoUrl, credentialsId: 'github-credential']
            ]
          ])
        }
      }
    }

    stage('Build Docker Image') {
      steps {
        script {
          sh 'docker build --platform linux/amd64 -t sikoding20/${IMAGE_NAME}:${BUILD_NUMBER} .'
        }
      }
    }

    stage('Push Docker Image') {
      steps {
        script {
          sh 'docker push sikoding20/${IMAGE_NAME}:${BUILD_NUMBER}'
        }
      }
    }

    stage('Update docker-compose.yaml') {
      steps {
        script {
          def runNumber = currentBuild.number
          sh "sed -i 's|image: sikoding20/${IMAGE_NAME}:[0-9]\\+|image: sikoding20/${IMAGE_NAME}:${runNumber}|' docker-compose.yaml"
        }
      }
    }

    stage('Commit and Push Changes') {
      steps {
        script {
          sh """
          git config --global user.name 'Jenkins CI'
          git config --global user.email 'jenkins@company.com'
          git remote set-url origin https://${GITHUB_CREDENTIALS_USR}:${GITHUB_CREDENTIALS_PSW}@github.com/bwa-project-example/payment-service.git
          git add docker-compose.yaml
          git commit -m 'Update image version to ${TARGET_BRANCH}-${currentBuild.number} [skip ci]' || echo 'No changes to commit'
          git pull origin ${TARGET_BRANCH} --rebase
          git push origin HEAD:${TARGET_BRANCH}
          """
        }
      }
    }
  }
}
