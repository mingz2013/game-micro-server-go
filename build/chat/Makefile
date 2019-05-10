help:
	@echo "Makefile help"

clean:
	rm chat-room-manager-go -f

chat-room-manager-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: chat-room-manager-go
	docker build -t mingz2013/chat-room-manager-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/chat-room-manager-go


remove-container:
	docker rm chat-room-manager-go

run:remove-container
	docker run -d --link redis-mq:redis-mq --name chat-room-manager-go -it mingz2013/chat-room-manager-go:latest

logs:
	docker logs web-proxy-server-go

.PYONY: help, commit-docker, docker-image, chat-room-manager-go, run, remove-container, logs

