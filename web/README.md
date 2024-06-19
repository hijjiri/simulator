### 扱ったコマンド一覧
docker build -t simulator-web:latest .
docker stop web
docker rm web
docker run -d --name web -p 3030:80 simulator-web:latest

# Heroku
cd /Users/saitoshosuke/go/src/github.com/hijjiri/simulator/web
heroku create simulator-web
heroku stack:set container -a simulator-web
docker build -t registry.heroku.com/simulator-web/web --platform linux/amd64 .
docker push registry.heroku.com/simulator-web/web
heroku container:release web -a simulator-web
