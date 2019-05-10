help:
	@echo "Makefile help"

clean:
	rm connector-server-go -f

connector-server-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: connector-server-go
	docker build -t mingz2013/connector-server-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/connector-server-go


remove-container:
	docker rm connector-server-go


run: remove-container
	docker run -d --link redis-mq:redis-mq --name connector-server-go -it mingz2013/connector-server-go:latest

logs:
	docker logs connector-server-go


.PYONY: help, commit-docker, docker-image, connector-server-go, run, remove-container, logs

