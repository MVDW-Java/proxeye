pipeline {
    agent any

    options {
        timeout(time: 15, unit: 'MINUTES')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Docker Deploy') {
            environment {
                PROXEYE_PORT = 8080
                GELBOORU_API_KEY = credentials('GELBOORU_API_KEY')
                GELBOORU_USER_ID = credentials('GELBOORU_USER_ID')
            }
            steps {
                sh "docker compose down --remove-orphans"
                sh "docker compose build --no-cache"
                sh "docker compose up -d"
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
