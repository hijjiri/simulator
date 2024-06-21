# simulator

# heroku

docker build -t registry.heroku.com/simulator-core/web --platform linux/amd64 .
docker build -t registry.heroku.com/simulator-web/web --platform linux/amd64 .
docker build -t registry.heroku.com/simulator-envoy/web --platform linux/amd64 .

docker push registry.heroku.com/simulator-core/web
docker push registry.heroku.com/simulator-web/web
docker push registry.heroku.com/simulator-envoy/web

heroku container:release web -a simulator-core
heroku container:release web -a simulator-web
heroku container:release web -a simulator-envoy

heroku open -a simulator-core
heroku open -a simulator-web
heroku open -a simulator-envoy

heroku info -a simulator-core
heroku info -a simulator-web
heroku info -a simulator-envoy

heroku logs --tail -a simulator-core
heroku logs --tail -a simulator-web
heroku logs --tail -a simulator-envoy

<!-- Herokuでは、アプリケーションはHerokuが提供するポートにバインドする必要があります。 -->
heroku config -a simulator-core
heroku config -a simulator-web
heroku config -a simulator-envoy

heroku config:set REACT_APP_GRPC_API_URL=https://simulator-core.herokuapp.com -a simulator-web



# 公開までのプロセス（議事録）

リポジトリの内容を直接解析するための情報を取得できませんでしたが、基本的なアプローチについてお答えします。

### 公開リポジトリの構造と懸念点

#### core
- **内容**: gRPCサーバーの実装。
- **懸念点**: サーバーのホスト名やポートが固定されている場合、Herokuでの動的な環境に対応できない可能性がある。
- **対策**: ホスト名やポートを環境変数で管理し、動的に設定できるようにする。

#### envoy
- **内容**: Envoyプロキシの設定ファイル。
- **懸念点**: 設定ファイルが固定されている場合、Herokuでの動的な環境に対応できない可能性がある。
- **対策**: Envoyの設定ファイルに環境変数を使用して、動的に設定を変更できるようにする。

#### web
- **内容**: フロントエンドのWebアプリケーション。
- **懸念点**: APIのエンドポイントやサーバーのホスト名が固定されている場合、Herokuでの動的な環境に対応できない可能性がある。
- **対策**: APIのエンドポイントやサーバーのホスト名を環境変数で管理し、動的に設定できるようにする。

### 具体的な課題と対応策

1. **環境変数の利用**
   - Herokuでは`config vars`を使用して環境変数を設定する。
   - コード内で環境変数を読み込み、設定に反映する。

2. **動的なホスト名とポートの対応**
   - gRPCサーバー、Envoyプロキシ、Webアプリケーションすべてでホスト名とポートを環境変数で管理する。
   - Herokuの`app.json`や`Procfile`で環境変数を設定する。

3. **設定ファイルのテンプレート化**
   - 必要に応じて設定ファイル（例：`envoy.yaml`）をテンプレート化し、デプロイ時に環境変数を置換するスクリプトを用意する。

これらの対策を講じることで、Herokuの動的な環境に適応し、安定した運用が可能になるでしょう。具体的な実装方法や設定について、さらに詳細な情報が必要な場合は、こちらに教えてください。

Herokuでの懸念点と対策
動的IPアドレス対応:

環境変数: Herokuの環境変数を利用して、動的に変化するIPアドレスやホスト名に対応。
設定方法:
sh
コードをコピーする
heroku config:set GRPC_SERVER_URL=yourapp.herokuapp.com:50051
Envoyの設定:

テンプレート利用: envoy.yaml.templateを使用し、envsubstで設定を生成。
Webアプリケーションのエンドポイント設定:

環境変数利用:
js
コードをコピーする
const apiEndpoint = process.env.REACT_APP_GRPC_API_URL || 'http://localhost:5000';

Herokuでは`docker-compose.yml`を直接使用できないため、Heroku向けに各サービスを個別にデプロイする必要があります。以下のステップで設定を行います。

### 1. **`Procfile`の作成**

Herokuで各サービスを定義するために`Procfile`を作成します。

```plaintext
web: cd web && npm start
core: cd core && ./core-server
envoy: ./start_envoy.sh
```

### 2. **`start_envoy.sh`スクリプトの作成**

`start_envoy.sh`スクリプトを作成し、Envoyを起動します。

```sh
#!/bin/sh

envsubst < /etc/envoy/envoy.yaml.template > /etc/envoy/envoy.yaml
/usr/local/bin/envoy -c /etc/envoy/envoy.yaml
```

### 3. **`app.json`の作成**

Herokuで環境変数を設定するために`app.json`を作成します。

```json
{
  "name": "simulator-app",
  "scripts": {},
  "env": {
    "GRPC_SERVER_URL": {
      "required": true
    },
    "REACT_APP_API_ENDPOINT": {
      "required": true
    }
  }
}
```

### 4. **環境変数の設定**

HerokuのダッシュボードまたはCLIから環境変数を設定します。

```sh
heroku config:set GRPC_SERVER_URL=yourapp.herokuapp.com:50051
heroku config:set REACT_APP_API_ENDPOINT=yourapp.herokuapp.com:8080
```

### 5. **Dockerfileの修正**

各サービスに対応するDockerfileを修正します。

#### `core/Dockerfile`

```Dockerfile
FROM golang:1.16

WORKDIR /app
COPY . .

RUN go build -o core-server ./cmd/core

CMD ["./core-server"]
```

#### `web/Dockerfile`

```Dockerfile
FROM node:14

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

CMD ["npm", "start"]
```

#### `envoy/Dockerfile`

```Dockerfile
FROM envoyproxy/envoy:v1.18.3

COPY envoy.yaml.template /etc/envoy/envoy.yaml.template
COPY start_envoy.sh /start_envoy.sh
RUN chmod +x /start_envoy.sh

CMD ["/start_envoy.sh"]
```

### 6. **Herokuへのデプロイ**

以下のコマンドでアプリケーションをHerokuにデプロイします。

```sh
git add .
git commit -m "Heroku deployment setup"
heroku create
git push heroku main
```

### 7. **Herokuでアプリケーションを起動**

Heroku上でアプリケーションが正しく起動するかを確認します。

```sh
heroku ps:scale web=1 core=1 envoy=1
```

これで、Herokuでのデプロイメントが完了し、各サービスが個別に起動されるようになります。
