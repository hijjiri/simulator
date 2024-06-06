# Dockerイメージのビルド
.PHONY: build
build:
	docker build -t simulator-web:latest .

# Dockerコンテナの停止と削除
.PHONY: stop
stop:
	@if docker ps -a | grep web; then docker stop web && docker rm web; fi

# Dockerコンテナの実行
.PHONY: run
run:
	docker run -d --name web -p 3000:80 simulator-web:latest

# Dockerコンテナの再ビルドと再実行
.PHONY: restart
restart: build run
	@echo "Docker container web has been restarted."
