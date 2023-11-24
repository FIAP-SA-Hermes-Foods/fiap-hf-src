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

        GPG_SECRET_KEY = credentials("GPG_SECRET_KEY")
        GPG_OWNER_TRUST = credentials("GPG_OWNER_TRUST")
        GPG_PASSWORD = credentials("GPG_SECRET_PASSWORD")
    }

    stages { 
        stage('Logging into AWS ECR') {
            steps {
                script {
                    sh """aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com"""
                }
            }
        }

        stage('Import secrets from git-secrets') { 
            steps {
                sh """gpg --batch --import ${GPG_SECRET_KEY}"""
                sh """gpg --import-ownertrust ${GPG_OWNER_TRUST}"""
            }
        }

        stage('Create .env') {
            steps {
                sh """git secret reveal -p '${GPG_PASSWORD}'"""
                sh """git secret cat .env > .env"""
            }
        } 

        stage('Create docker network') {
            steps {
                sh './infrastructure/scripts/docker-network.sh'    
            }
        }

        stage('Building images') {
            steps{
                script {
                    sh """docker build -f infrastructure/docker/go/Dockerfile -t ${IMAGE_API_NAME}:${IMAGE_TAG} ."""
                    sh """docker build -f infrastructure/docker/postgres/Dockerfile -t ${IMAGE_POSTGRES_NAME}:${IMAGE_TAG} ."""
                    sh """docker build -f infrastructure/docker/swagger/Dockerfile -t ${IMAGE_SWAGGER_NAME}:${IMAGE_TAG} ."""
                    sh 'docker images'
                }
            }
        }

        stage('Tagging images') {
            steps{
                script {
                    sh """docker tag ${IMAGE_API_NAME}:${IMAGE_TAG} ${REPOSITORY_API_URL}:$IMAGE_TAG"""
                    sh """docker tag ${IMAGE_POSTGRES_NAME}:${IMAGE_TAG} ${REPOSITORY_POSTGRES_URL}:$IMAGE_TAG"""
                    sh """docker tag ${IMAGE_SWAGGER_NAME}:${IMAGE_TAG} ${REPOSITORY_SWAGGER_URL}:$IMAGE_TAG"""
                }
            }
        }

        stage('Pushing to ECR') {
            steps{  
                script {
                    sh """docker push ${REPOSITORY_API_URL}:${IMAGE_TAG}"""
                    sh """docker push ${REPOSITORY_POSTGRES_URL}:${IMAGE_TAG}"""
                    sh """docker push ${REPOSITORY_SWAGGER_URL}:${IMAGE_TAG}"""
                }
            }
        }

        stage('Kubernetes setup') {
            steps {
                sh './infrastructure/scripts/kubernetes-config.sh'
            }
        }

        stage('set environments') {
            steps {
                script {
                    sh '''#!/bin/bash
                        if [ -f .env ]; then
                            cp -f .env $HOME/envs

                            if [ -f $HOME/envs/.env.export ]; then 
                                rm -f $HOME/envs/.env.export
                            fi
          
                            cat $HOME/envs/.env | while read LINE; do
                                if [[ $LINE == \\#* ]]; then
                                    continue
                                fi
                                export $LINE
                                echo "export $LINE" >> $HOME/envs/.env.export
                            done
                        fi
                    '''
                } 
            }
        }
        
        stage('sourcing...') {
            steps {
                script{ 
                    sh '. $HOME/envs/.env.export'
                }
            }

        stage('Deploy at k8s') {
            steps { 
                script {
                    sh 'kubectl apply -f ./etc/kubernetes/config/postgres.yaml'
                    sh 'kubectl apply -f ./etc/kubernetes/deployment/app.yaml'
                    sh 'kubectl apply -f ./etc/kubernetes/deployment/postgres.yaml'
                    sh 'kubectl apply -f ./etc/kubernetes/deployment/swagger.yaml'
                }
            }
        }
    }
}
