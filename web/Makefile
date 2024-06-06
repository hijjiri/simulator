# Dockerイメージのビルド
.PHONY: docker-build
docker-build:
	docker build -t hijjiri-web:latest .

# Dockerコンテナの停止と削除
.PHONY: docker-stop
docker-stop:
	@if docker ps -a | grep hijjiri-web; then docker stop hijjiri-web && docker rm hijjiri-web; fi

# Dockerコンテナの実行
.PHONY: docker-run
docker-run:
	docker run -d --name hijjiri-web -p 3030:80 hijjiri-web:latest

# Dockerコンテナの再ビルドと再実行
.PHONY: docker-restart
docker-restart: docker-build docker-run
	@echo "Docker container hijjiri-web has been restarted."
