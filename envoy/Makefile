.PHONY: build run up down

build:
	docker run -d --name envoy -p 8080:8080 -v $(PWD)/envoy.yaml:/etc/envoy/envoy.yaml envoyproxy/envoy:v1.18.3

run:
	docker restart envoy

up:
	docker-compose up -d

down:
	docker-compose down
