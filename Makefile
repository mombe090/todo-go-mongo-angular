build:
	docker-compose build 

run:
	docker-compose up -d

log:
	docker-compose logs -f todo-backend

all: build run

#https://nikgrozev.com/2018/10/12/python-api-with-flask-and-flask-restplus/