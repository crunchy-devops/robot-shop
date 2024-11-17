pipeline {
    agent any
    stages {
        stage('Checkout Code') {
            steps {
                checkout scmGit(tags: [[name: 'v1']],
                    userRemoteConfigs: [[url: 'https://github.com/crunchy-devops/robot-shop.git']])
                //git branch:'master', url: 'https://github.com/crunchy-devops/robot-shop.git'
                //git 'https://github.com/crunchy-devops/robot-shop.git'
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
        stage('create build report') {
            steps{
                sh "docker images --all --filter=reference='robotshop/*:${env.GIT_COMMIT_HASH}' --format '{{.Repository}}\t{{.Size}}' >/bitnami/jenkins/home/checkimages/build/build-${env.BUILD_NUMBER}"
            }
        }
        stage('check docker images size') {
            steps{
                sh 'python3 /bitnami/jenkins/home/checkimages/check-docker-images.py'
            }
        }
        stage('alert docker images size') {
            steps {
                script {
                    // Read the logs of the current build
                    def log = currentBuild.rawBuild.getLog(1000) // Reads up to 1000 lines
                    def pattern = /Value:\s*([0-9]*\.[0-9]+)/

                    // Iterate over each line in the log
                    log.each { line ->
                        def matcher = (line =~ pattern)
                        if (matcher.find()) {
                            // Extract the value and convert it to an integer

                            def value = matcher.group(1).toFloat()

                            // Check if the value is out of the specified limits, 25%
                            if (value > 25.00 ) {

                                echo "ALERT: Value ${value} is out of limits!"
                                // You can mark the build as unstable, or fail it based on requirements
                                currentBuild.result = 'UNSTABLE'
                            }
                        }
                    }
                }
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
        stage('install docker'){
            steps {
                    ansibleTower jobTemplate: 'install docker', jobType: 'run', throwExceptionWhenFail: false, towerCredentialsId: 'ansiblelogin', towerLogLevel: 'full', towerServer: 'awx'
                }
        }
   }
}