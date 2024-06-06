PROTO_DIR=./proto
OUTPUT_DIR=./go/grpc-server
JS_DIR=./src/assets/_proto
GOOGLEAPIS_DIR=/Users/saitoshosuke/go/src/github.com/googleapis

# プロトコルバッファのコンパイル
.PHONY: proto

proto:
	for file in $(PROTO_DIR)/*.proto; do \
		filename=$$file; \
		dirname=$$(basename $$file .proto); \
		mkdir -p "$(OUTPUT_DIR)/$${dirname}"; \
		protoc --js_out=import_style=commonjs,binary:$(JS_DIR) \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:$(JS_DIR) \
		--go_out=paths=source_relative:$(OUTPUT_DIR)/$${dirname} \
		--go-grpc_out=paths=source_relative:$(OUTPUT_DIR)/$${dirname} \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:$(OUTPUT_DIR)/$${dirname} \
		--proto_path=$(PROTO_DIR) \
		--proto_path=$(GOOGLEAPIS_DIR) \
		$$file; \
	done

# Dockerイメージのビルド
.PHONY: build
build:
	docker build -t simulator-core:latest .

# Dockerコンテナの停止と削除
.PHONY: stop
stop:
	@if docker ps -a | grep core; then docker stop core && docker rm core; fi

# Dockerコンテナの実行
.PHONY: run
run:
	docker run -d --name core -p 50051:50051 simulator-core:latest

# Dockerコンテナの再ビルドと再実行
.PHONY: restart
restart: build run
	@echo "Docker container core has been restarted."
