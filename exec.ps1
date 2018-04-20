docker image build --no-cache -t auth0-golang-kube-app .
docker container run -it --rm --env-file env.dev -p 3000:3000 auth0-golang-kube-app
