### 扱ったコマンド一覧
docker build -t simulator-web:latest .
docker stop web
docker rm web
docker run -d --name web -p 3030:80 simulator-web:latest
