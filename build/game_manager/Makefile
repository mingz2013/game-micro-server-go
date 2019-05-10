help:
	@echo "Makefile help"

clean:
	rm game-manager-server-go -f

game-manager-server-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: game-manager-server-go
	docker build -t mingz2013/game-manager-server-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/game-manager-server-go


remove-container:
	docker rm game-manager-server-go


run: remove-container
	docker run -d --link redis-mq:redis-mq --name game-manager-server-go -it mingz2013/game-manager-server-go:latest

logs:
	docker logs game-manager-server-go

.PYONY: help, commit-docker, docker-image, game-manager-server-go, run, remove-container, logs

