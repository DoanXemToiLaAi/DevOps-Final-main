pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'whyyouask/devops-book'
        DOCKER_TAG = 'latest'
        TELEGRAM_BOT_TOKEN = '8095963368:AAEkxTTHPG6BjvawtWEa189CF6Gf9NGrlBA'
        TELEGRAM_CHAT_ID = '5603549047'
    }

    stages {
         stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/DoanXemToiLaAi/DevOps-Final-main'
            }
        }
                  stage('Build Docker Image') {
            steps {
                script {
                    echo 'Building Docker image for linux/amd64 platform...'
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}", "--platform linux/amd64 .")
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                script {
                    echo 'Clearing server_golang-related images and containers...'
                    sh '''
                        docker container stop server-golang || echo "No container named server-golang to stop"
                        docker container rm server-golang || echo "No container named server-golang to remove"
                        docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG} || echo "No image ${DOCKER_IMAGE}:${DOCKER_TAG} to remove"
                    '''
                    
                    echo 'Deploying to DEV environment...'
                    sh '''
                        docker image pull ${DOCKER_IMAGE}:${DOCKER_TAG}
                        docker network create dev || echo "Network already exists"
                        docker container run -d --rm --name server-golang -p 3005:3005 --network dev ${DOCKER_IMAGE}:${DOCKER_TAG}
                    '''
                }
            }
        }
    }

      post {
        always {
            cleanWs()
        }

        success {
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}
def sendTelegramMessage(String message) {
    sh """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
    -d chat_id=${TELEGRAM_CHAT_ID} \
    -d text="${message}"
    """
}