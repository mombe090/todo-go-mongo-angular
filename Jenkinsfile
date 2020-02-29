pipeline {
    agent { 
            node {
                    label 'production'
                } 
        }
    stages {
        
        stage('Checkout') {
            
            steps {
                checkout scm
                sh 'ls -l'
            }
        }

        stage('Build') {
            
            
            steps {
                
                sh  """
                        docker build  -t todo_mongo_go:1.0.0 .
                    """
            }
        }

        stage('Build and Deploy') {
            
            
            steps {
                sh  """
                    docker rm -f todo_mongo_training ||  true
                    docker run -dit  --label traefik.docker.network=proxy_https_default  --label traefik.frontend.rule=Host:todogo.guineeapps.com  --label traefik.port=8090  --label traefik.backend=todo_mongo_training --name todo_mongo_training --network  proxy_https_default  todo_mongo_go:1.0.0
               
                    """
               
            }
        }

    }
}