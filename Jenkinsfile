pipeline {
    agent any
    environment {
        AWS_ACCOUNT_ID = credentials('AWS_ACCOUNT_ID')
        AWS_DEFAULT_REGION = credentials('AWS_DEFAULT_REGION')
        IMAGE_API_NAME = credentials('IMAGE_API_NAME')
        IMAGE_POSTGRES_NAME = credentials('IMAGE_POSTGRES_NAME')
        IMAGE_SWAGGER_NAME = credentials('IMAGE_SWAGGER_NAME')
        IMAGE_TAG= "latest"
        REPOSITORY_API_URL = credentials('ECR_API_URL')
        REPOSITORY_POSTGRES_URL = credentials('ECR_POSTGRES_URL')
        REPOSITORY_SWAGGER_URL = credentials('ECR_SWAGGER_URL')
        HF_GITHUB_PROJECT_URL = credentials('HF_GITHUB_PROJECT_URL')
    }

    stages { 
        stage('Logging into AWS ECR') {
            steps {
                script {
                    sh """aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com"""
                }
            }
        }

        stage('Rename .env') {
            steps {
                sh 'mv .env.example .env'    
            }
        }

        // Building Docker images
        stage('Building images') {
            steps{
                script {
                    sh 'make run-build-d'
                }
            }
        }

        // Uploading Docker images into AWS ECR
        stage('Pushing to ECR') {
            steps{  
                script {
                    sh """docker tag ${IMAGE_API_NAME}:${IMAGE_TAG} ${REPOSITORY_API_URL}:$IMAGE_TAG"""
                        sh """docker push ${REPOSITORY_API_URL}:${IMAGE_TAG}"""
                        sh """docker tag ${IMAGE_POSTGRES_NAME}:${IMAGE_TAG} ${REPOSITORY_POSTGRES_URL}:$IMAGE_TAG"""
                        sh """docker push ${REPOSITORY_POSTGRES_URL}:${IMAGE_TAG}"""
                        sh """docker tag ${IMAGE_SWAGGER_NAME}:${IMAGE_TAG} ${REPOSITORY_SWAGGER_URL}:$IMAGE_TAG"""
                        sh """docker push ${REPOSITORY_SWAGGER_URL}:${IMAGE_TAG}"""
                }
            }
        }
    }
}
