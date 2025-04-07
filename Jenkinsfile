pipeline {
    agent any

    environment{
        IMAGE_NAME = 'sicei'
        CONTAINER_NAME = 'sicei_api'
    }
    stages{
        stage('Build'){
            steps{
                sh 'docker build -t ${IMAGE_NAME}:${BUILD_ID} .'

                sh 'docker stop ${CONTAINER_NAME} || true' 
                sh 'docker rm ${CONTAINER_NAME} || true '
            }
        }
        stage('Deploy'){
            steps{
                sh 'docker run -p 80:80 --name ${CONTAINER_NAME} -d ${IMAGE_NAME}:${BUILD_ID}'
            }
        }
    }
}