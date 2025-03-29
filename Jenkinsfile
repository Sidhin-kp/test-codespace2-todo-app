pipeline {
    agent any

    environment {
        DOCKER_IMAGE_BACKEND = "todoapp-backend"
        DOCKER_IMAGE_FRONTEND = "todoapp-frontend"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/Sidhin-kp/test-codespace2-todo-app.git'
            }
        }

        stage('Build Backend Image') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE_BACKEND ./backend'
            }
        }

        stage('Build Frontend Image') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE_FRONTEND ./frontend'
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
