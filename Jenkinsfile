pipeline {
    agent any

    environment {
        DOCKER_IMAGE_BACKEND = "todoapp-backend"
        DOCKER_IMAGE_FRONTEND = "todoapp-frontend"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main',
                    credentialsId: 'github-creds',
                    url: 'https://github.com/Sidhin-kp/test-codespace2-todo-app.git'
            }
        }

        stage('Build Backend Image') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE_BACKEND -f backend/Dockerfile.backend backend/'
            }
        }

        stage('Build Frontend Image') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE_FRONTEND -f frontend/Dockerfile.frontend frontend/'
            }
        }

        stage('Start Services') {
            steps {
                sh 'docker-compose down'
                sh 'docker-compose up -d'
            }
        }

        stage('Run Tests') {
            steps {
                sh 'docker-compose -f docker-compose.test.yml up --abort-on-container-exit'
            }
        }

        stage('Cleanup') {
            steps {
                sh 'docker system prune -f'
            }
        }
    }
}
