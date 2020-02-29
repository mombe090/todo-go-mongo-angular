# **TODO RESTFULL WITH GOLANG & ANGULAR**

## Getting Started

To test this project, we have the following way

* With Docker, Docker-Compose ou Docker-Swarm
* Building natively using the go cli-tools
* --- comming on savoirguinee


### Prerequisites

To install and run this project, you need the following softs and tools

* golang SKD installed if you want tested it as developer. go download the latest stable SDK depending on your OS https://golang.org/dl/ and a mongodb server up and runing
* docker if you want first testing the application without coding anything and installing a mongodb server
* we will give an url for direct test on https://mooc.savoirguinee.com/


### Installing

To run this application, you need first to clone the code locally on your computer using git or download it as zip on github https://github.com/mombe090/todo-go-mongo-angular

```
git clone https://github.com/mombe090/todo-go-mongo-angular.git
```
Then

```
cd todo-go-angular
```

If you opt for Docker, please use the following step


## UP AND RUN WITH DOCKER

* With Docker
    - ```
            cd backend
      ```
    - update the MONGOHOST and MONGOPORT by yours in the .env file
    - ```
            docker build -t name-of-your-image .
      ```
    - ```
            docker run -n name-of-your-container -d -p 8090:8090 name-of-your-image
      ```
      Here we go, you can test it on http://localhost:8090/todos
      
* With Docker-Compose
    - update the MONGOHOST ***mongo*** if you've changed it in the .env file
    - we assume you are in the tod folder
     - ```
             docker-compose up --build -d 
       ```
        Here we go, you can test it on http://localhost:8090/todos

* With Docker-Swarm
     to run the app in the swarm, we have init the swarm first
     - ```
             docker swarn init  
       ```
     - ```
             docker stack deploy -c docker-stack.ym todo   
       ```  
        Here we go, you can test it on http://localhost:8090/todos


## UP AND RUN WITH GOLANG CLI

* How to build the application using the cli
    - update the MONGOHOST and MONGOPORT by yours in the .env file
    - ```
            cd backend  
      ``` 
    - ```
            go build -o todo    
      ```
      add .exe if you are in Windows
     - ```
            ./todo    
       ```
     Here we go, you can test it on http://localhost:8090

## Built With

* [MongoDB](https://www.mongodb.com/) - The database for modern applications
* [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Gin-GONIC](https://gin-gonic.com/) - Gin is a HTTP web framework written in Go (Golang).

## Contributors 

[Mombe090](https://github.com/mombe090) &
[Rougeo](https://github.com/rougeo)

## Authors

* **Mamadou Yaya DIALLO** - *Initial work* - [Devscom](https://twitter.com/DevscomGn)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

