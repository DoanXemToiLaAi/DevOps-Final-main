pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'whyyouask/devops-book'
        DOCKER_TAG = 'latest'
        TELEGRAM_BOT_TOKEN = '8046339515:AAEIJtDScmfi0ExQFrk4ATFCKfIJYsFVdJY'
        TELEGRAM_CHAT_ID = '-1002515055682'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/DoanXemToiLaAi/DevOps-Final-main.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                bat "docker build -t %DOCKER_IMAGE%:%DOCKER_TAG% --platform linux/amd64 ."
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    bat '''
                    echo ===== Logging in to Docker Hub =====
                    echo %DOCKER_PASS% | docker login -u %DOCKER_USER% --password-stdin

                    echo ===== Pushing Docker Image =====
                    docker push %DOCKER_IMAGE%:%DOCKER_TAG%
                    '''
                }
            }
        }

        stage('Deploy Container') {
            steps {
                bat '''
                echo ===== Stopping and removing old container =====
                docker stop devops-book || echo Not running
                docker rm devops-book || echo Not found

                echo ===== Removing old image =====
                docker image rm %DOCKER_IMAGE%:%DOCKER_TAG% || echo No image to remove

                echo ===== Pulling latest image =====
                docker pull %DOCKER_IMAGE%:%DOCKER_TAG%

                echo ===== Creating network (if needed) =====
                docker network create dev || echo Already exists

                echo ===== Starting new container =====
                docker run -d --name devops-book -p 4200:4200 --network dev %DOCKER_IMAGE%:%DOCKER_TAG%
                '''
            }
        }
    }

    post {
        success {
            script {
                sendTelegramMessage("Build #${env.BUILD_NUMBER} completed successfully.")
            }
        }

        failure {
            script {
                sendTelegramMessage("Build #${env.BUILD_NUMBER} failed. Please check Jenkins logs.")
            }
        }
    }
}

def sendTelegramMessage(String message) {
    bat """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage ^
    -d chat_id=${TELEGRAM_CHAT_ID} ^
    -d text="${message}"
    """
}
