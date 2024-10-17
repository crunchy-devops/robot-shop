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
                sh 'docker-compose down && docker-compose up -d'
            }
        }
        stage('Tag images and push them to nexus') {
            steps{
                script {
                          def nexus = "nexus:30999"
                          def output = sh(script: "docker images --all   --filter=reference='robotshop/*:${env.GIT_COMMIT_HASH}'  --format '{{.Repository}}'", returnStdout: true).trim()
                          def imageArray = output.split("\n")

                          for (image in imageArray) {
                            echo "Processing image: ${image}"
                            def inter = image.split("/")
                            echo " ${inter[0]} and ${inter[1]}"
                            def nexusimage = nexus + "/" + inter[1] + ":" + env.GIT_COMMIT_HASH
                            echo "${nexusimage}"
                            def imageTag = image + ":" + env.GIT_COMMIT_HASH
                            def tag = sh(script: "docker tag ${imageTag} ${nexusimage}", returnStdout: true).trim()
                            echo "$tag"
                            withCredentials([usernamePassword(credentialsId: 'nexuslogin', passwordVariable: 'PASSWORD', usernameVariable: 'USERNAME')]) {
                               def push = sh(script: "docker login ${nexus} -u ${USERNAME} -p ${PASSWORD} && docker push ${nexusimage}" , returnStdout: true).trim()
                            }
                       }
                    }
                }
        }
    }
}