FROM golang:1.14-alpine as build

WORKDIR /go/reaction

COPY src/app ./src/app
COPY go.mod .
COPY go.sum .
COPY .realize.yml .

RUN apk add --no-cache git \
  && go get github.com/oxequa/realize

WORKDIR /go/reaction/src/app

RUN go build -o app

FROM alpine as prod

WORKDIR /app

COPY --from=build /go/reaction/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /reaction/app

CMD ["./app"]