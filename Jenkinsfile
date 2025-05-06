pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'whyyouask/devops-boks'
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
                bat '''
                echo === CHECKING DIRECTORY ===
                cd
                dir
                echo === BUILDING DOCKER IMAGE ===
                docker build -t %DOCKER_IMAGE%:%DOCKER_TAG% .
                '''
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    bat '''
                    echo === LOGGING IN TO DOCKER HUB ===
                    echo %DOCKER_PASS% | docker login -u %DOCKER_USER% --password-stdin

                    echo === PUSHING IMAGE ===
                    docker push %DOCKER_IMAGE%:%DOCKER_TAG%
                    '''
                }
            }
        }

        stage('Deploy Container') {
            steps {
                bat '''
                echo === STOP AND REMOVE OLD CONTAINER ===
                docker stop devops-boks || echo Container not running
                docker rm devops-boks || echo Container not found

                echo === REMOVE OLD IMAGE ===
                docker rmi %DOCKER_IMAGE%:%DOCKER_TAG% || echo Image not found

                echo === PULL LATEST IMAGE ===
                docker pull %DOCKER_IMAGE%:%DOCKER_TAG%

                echo === CREATE NETWORK ===
                docker network create dev || echo Network already exists

                echo === STARTING NEW CONTAINER ===
                docker run -d --name devops-boks -p 4200:4200 --network dev %DOCKER_IMAGE%:%DOCKER_TAG%
                '''
            }
        }
    }

    post {
        success {
            script {
                sendTelegramMessage("✅ Build #${env.BUILD_NUMBER} was successful!")
            }
        }
        failure {
            script {
                sendTelegramMessage("❌ Build #${env.BUILD_NUMBER} failed. Please check Jenkins logs.")
            }
        }
    }
}

def sendTelegramMessage(String message) {
    bat """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage ^
    -d chat_id=${TELEGRAM_CHAT_ID} ^
    -d text=\"${message}\"
    """
}
