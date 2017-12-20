# Auth0 - Go Token Web App

This sample app has been modified to produce an `id_token` and `refresh_token` to  stdout in the console once you've authenticated through your web browser


### Prerequisites

You'll need an [Auth0](https://auth0.com/) account and web [client](https://auth0.com/docs/clients)

Auth0 Jargon:
* Auth0 Regular WebApp client
* Your client must be authorised in both the management and extensions API
* The extensions API should have `read:user_idp_tokens` scope and the __Allow Offline Access__ activated


### Usage

Assuming you have set up a client in [Auth0](https://auth0.com/)?

You'll need to retrieve values from your Auth0 client for the environment variables listed in the `.env` file

Once done, simply run the `exec.sh` script and browse to [http://localhost:3000](http://localhost:3000) where you can authenticate

Your `id_token` and `refresh_token` output can be found in the console you ran the `exec.sh` from
