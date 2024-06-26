# core

# .envrcを適用させたら以下のコマンドを実行
```
//出力されればOK！
which protoc
which protoc-gen-go-grpc
which protoc-gen-grpc-gateway
which go
which protoc-gen-go
which protoc-gen-grpc-web
which protoc-gen-js
```

#　インストールと実行ファイル
```
sudo apt install protobuf-compiler
protoc --version
```

```
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH=$PATH:$(go env GOPATH)/bin
```

```
wget https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v2.19.0/protoc-gen-grpc-gateway-v2.19.0-linux-x86_64
chmod +x protoc-gen-grpc-gateway-v2.19.0-linux-x86_64
sudo mv protoc-gen-grpc-gateway-v2.19.0-linux-x86_64 /usr/local/bin/protoc-gen-grpc-gateway
```

```
go get google.golang.org/protobuf/cmd/protoc-gen-go
//実行:
protoc --go_out=./go/grpc-server/${dirname} $file
```

```
brew install protoc-gen-grpc-web
実行:
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/assets/_proto $file
```

```
yarn global add grpc-web
実行：
protoc --js_out=import_style=commonjs,binary:./src/assets/_proto $file
```


# yarn手順
```
yarn global add protoc-gen-js
yarn global bin
export PATH=$PATH:/Users/saitoshosuke/.yarn/bin
source ~/.zshrc
```

# 扱ったコマンド一覧
docker build -t simulator-core:latest .
docker stop core
docker rm core
docker run -d --name core -p 50051:50051 simulator-core:latest

docker network ls
docker network rm my_network
docker network create my_network
docker network connect my_network core
docker network connect my_network web
docker network connect my_network envoy
docker network inspect my_network
<!-- coreのIP取得 -->
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' envoy-core-1

grpcurl -plaintext localhost:50051 list
grpcurl -plaintext localhost:50051 describe example.ExampleService
grpcurl -plaintext -import-path /Users/saitoshosuke/go/src/github.com/hijjiri/simulator/core/proto -proto example.proto -d '{"name": "world"}' localhost:50051 example.ExampleService/SayHello

# Heroku
heroku create simulator-core
heroku stack:set container -a simulator-core
docker build -t registry.heroku.com/simulator-core/web --platform linux/amd64 .
docker push registry.heroku.com/simulator-core/web
heroku container:release web -a simulator-core

# Heroku test
grpcurl -plaintext -d '{"name": "world"}' simulator-core.herokuapp.com:80 example.ExampleService/SayHello