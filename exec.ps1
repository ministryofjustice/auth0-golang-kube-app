docker build --no-cache -t auth0-golang-kube-app .
docker run -it --env-file .env -p 3000:3000 auth0-golang-kube-app