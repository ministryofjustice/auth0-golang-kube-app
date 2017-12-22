# Auth0 - Go Token Web App

This sample app has been modified to produce an `id_token` and `refresh_token` to stdout and your web browser once you've authenticated with [Auth0](https://auth0.com/)

My use case was for Kubernetes cluster auth, namely [Option 1](https://kubernetes.io/docs/admin/authentication/#using-kubectl) when configuring OpenID Connect for `kubectl`


### Prerequisites

You'll need an [Auth0](https://auth0.com/) account and web [client](https://auth0.com/docs/clients)

Auth0 Jargon:
* Auth0 Regular WebApp client
* Your client must be authorised in both the management and extensions API
* The extensions API should have `read:user_idp_tokens` scope and the __Allow Offline Access__ activated


__Required Environments Variables__

Values can be found in [Auth0](https://auth0.com/) client settings

```
AUTH0_CLIENT_ID=< client id >
AUTH0_DOMAIN=< domain >
AUTH0_CLIENT_SECRET=< client_secret >
AUTH0_CALLBACK_URL=http://localhost:3000/callback
AUTH0_AUDIENCE=urn:auth0-authz-api
```

### Usage

__With Docker__

Assuming you have set up a client in [Auth0](https://auth0.com/)?

You'll need to retrieve values from your Auth0 client for the environment variables listed in the `.env` file

Once done, simply run the `exec.sh/exec.ps1` script and browse to [http://localhost:3000](http://localhost:3000) where you can authenticate

Your `id_token` and `refresh_token` output can be copy&pasted from the console you ran the `exec.sh` from or your browser

See [refresh_token](https://auth0.com/docs/tokens/refresh-token/current) and [id_token](https://auth0.com/docs/tokens/id-token#overview) docs for more info


__or__

Use my container and supply your own `.env` file

`docker run -d --env-file .env -p 3000:3000 quay.io/djsd123/auth0-golang-kube-app`


__Without Docker__

* First set the environment variables defined above

* Run `go run main.go server.go`

* Browse to [http://localhost:3000](http://localhost:3000)
