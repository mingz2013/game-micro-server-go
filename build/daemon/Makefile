help:
	@echo "Makefile help"

clean:
	rm daemon-server-go -f

daemon-server-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: daemon-server-go
	docker build -t mingz2013/daemon-server-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/daemon-server-go



remove-container:
	docker rm daemon-server-go


run: remove-container
	docker run -d --link redis-mq:redis-mq --name daemon-server-go -it mingz2013/daemon-server-go:latest


logs:
	docker logs daemon-server-go


.PYONY: help, commit-docker, docker-image, daemon-server-go, run, remove-container, logs

