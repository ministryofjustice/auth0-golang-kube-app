FROM golang:1.8-alpine3.6

# install git
RUN apk --update add \
	git openssl \
	&& rm /var/cache/apk/*

WORKDIR /go/src

ADD . /go/src

ENV AUTH0_CLIENT_ID=oUb1V330oXKyMpTagAYDzWDY10U4ffWF
ENV AUTH0_DOMAIN=alpha-analytics-moj.eu.auth0.com
ENV AUTH0_CLIENT_SECRET=dXlUcehHp6M9lwVAOhIkJ7A_6IMhlQMEBp2Mhr5rUdo4K6NMSq9z6fdzP-yo9o9D
ENV AUTH0_CALLBACK_URL=http://localhost:3000/callback
ENV AUTH0_AUDIENCE=urn:auth0-authz-api

RUN go-wrapper download

CMD ["go", "run", "main.go", "server.go"]

EXPOSE 3000
