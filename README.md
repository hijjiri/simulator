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
