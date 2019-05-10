help:
	@echo "Makefile help"

clean:
	rm login-server-go -f

login-server-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: login-server-go
	docker build -t mingz2013/login-server-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/login-server-go



remove-container:
	docker rm login-server-go


run: remove-container
	docker run -d --link redis-mq:redis-mq --name login-server-go -it mingz2013/login-server-go:latest

logs:
	docker logs login-server-go

.PYONY: help, commit-docker, docker-image, login-server-go, run, remove-container, logs

