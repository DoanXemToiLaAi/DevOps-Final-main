pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'whyyouask/devops-book'
        DOCKER_TAG = 'latest'
        TELEGRAM_BOT_TOKEN = '8095963368:AAEkxTTHPG6BjvawtWEa189CF6Gf9NGrlBA'
        TELEGRAM_CHAT_ID = '-4715805728'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/DoanXemToiLaAi/DevOps-Final-main.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} --platform linux/amd64 ."
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push ${DOCKER_IMAGE}:${DOCKER_TAG}
                    '''
                }
            }
        }

        stage('Deploy Container') {
            steps {
                sh '''
                    docker stop devops-book || true
                    docker rm devops-book || true
                    docker network create dev || true
                    docker run -d --rm --name devops-book -p 4200:4200 --network dev ${DOCKER_IMAGE}:${DOCKER_TAG}
                '''
            }
        }
    }

    post {
        success {
            script {
                sendTelegramMessage("✅ Build #${env.BUILD_NUMBER} thành công trên Jenkins!")
            }
        }

        failure {
            script {
                sendTelegramMessage("❌ Build #${env.BUILD_NUMBER} thất bại.")
            }
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
