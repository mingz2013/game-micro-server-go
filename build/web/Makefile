help:
	@echo "Makefile help"

clean:
	rm web-proxy-server-go

web-proxy-server-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: web-proxy-server-go
	docker build -t mingz2013/web-proxy-server-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/web-proxy-server-go



remove-container:
	docker rm web-proxy-server-go


run: remove-container
	docker run -d -p 8006:8006 --link redis-mq:redis-mq --name web-proxy-server-go -it mingz2013/web-proxy-server-go:latest


logs:
	docker logs web-proxy-server-go


.PYONY: help, commit-docker, docker-image, web-proxy-server-go, run, remove-container, logs

