pipeline {
    agent any
    stages {
        stage('Checkout Code') {
            steps {
                git 'https://github.com/crunchy-devops/robot-shop.git'
                //checkout scm
            }
        }
        stage('Retrieve Git Commit Hash') {
            steps {
                script {
                    // Retrieve the Git commit hash using 'git rev-parse'
                    env.GIT_COMMIT_HASH = sh(script: 'git rev-parse HEAD | cut -c 1-12', returnStdout: true).trim()
                }
            }
        }
        stage('Display Git Commit Hash') {
            steps {
                // Display the retrieved Git commit hash using echo
                echo "The current Git commit hash is: ${env.GIT_COMMIT_HASH}"
            }
        }
        stage('Replace TAG in .env') {
            steps{
                // replace
                sh "sed -i -e 's/TAG=2.1.0/TAG=${env.GIT_COMMIT_HASH}/g' .env"
                sh 'cat .env'
            }
        }
        stage('Docker-compose build'){
            steps{
                // Docker
                sh 'docker-compose build'
            }
        }
        stage('Docker-compose up') {
            steps{
                sh 'docker-compose up -d'
            }
        }
    }
}