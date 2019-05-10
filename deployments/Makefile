help:
	@echo 'Makefile help                                                             '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make docker-compose-up                           创建并启动所有容器       '
	@echo '   make docker-compose-stop                         停止容器               '
	@echo '   make docker-compose-start                        启动                   '
	@echo '   make docker-compose-down                         删除                   '
	@echo '   make docker-compose-ps                           查看状态               '
	@echo '                                                                          '
	@echo '   make start-redis-single                           单独启动redis         '
	@echo '                                                                          '

docker-compose-up:
	docker-compose up -d

docker-compose-stop:
	docker-compose stop

docker-compose-start:
	docker-compose start

docker-compose-down:
	docker-compose down

docker-compose-ps:
	docker-compose ps


start-redis-single:
	docker run -d -p 6379:6379 --name redis-mq redis:latest


.PYONY: help, docker-compose-up, docker-compose-stop, docker-compose-start, docker-compose-down, docker-compose-ps, start-redis-single

