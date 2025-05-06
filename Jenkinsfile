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
                git branch: 'main', url: 'https://github.com/DoanXemToiLaAi/DevOps-Final-main.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    echo '🛠 Building Docker image...'
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}", "--platform linux/amd64 .")
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                script {
                    echo '♻️ Clearing previous deployment...'
                    sh '''
                        docker container stop server-golang || true
                        docker container rm server-golang || true
                        docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG} || true
                    '''

                    echo '🚀 Deploying new container...'
                    sh '''
                        docker network create dev || true
                        docker run -d --rm --name server-golang -p 4200:4200 --network dev ${DOCKER_IMAGE}:${DOCKER_TAG}
                    '''
                }
            }
        }
    }

    post {
        success {
            script {
                sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful!")
            }
        }
        failure {
            script {
                sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed.")
            }
        }
    }
}

def sendTelegramMessage(String message) {
    sh """
        curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \\
        -d chat_id=${TELEGRAM_CHAT_ID} \\
        -d text="${message}"
    """
}
