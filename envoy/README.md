# 扱ったコマンド一覧

docker run -d --name envoy -p 8080:8080 -v $(PWD)/envoy.yaml:/etc/envoy/envoy.yaml envoyproxy/envoy:v1.18.3

docker restart envoy

# docker-compose
docker-compose down
docker-compose up -d

# Jenkins
sudo lsof -i :8080 | grep LISTEN
sudo kill -9 <PID>
brew services restart jenkins-lts

# sshサーバーの設定関連
## macOSの場合はデフォルトでsshサーバーが入っているので、以下コマンドでオンに設定
sudo systemsetup -setremotelogin on

# heroku
heroku apps:destroy --app portfolio-simulator --confirm portfolio-simulator
heroku create portfolio-simulator
git remote rm heroku
heroku git:remote -a portfolio-simulator
heroku buildpacks:clear
heroku stack:set container
