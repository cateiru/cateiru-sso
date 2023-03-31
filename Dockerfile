FROM golang:alpine as builder

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /go/src

# OIDCのJWTに使用する公開鍵と秘密鍵を作成する
RUN apk add openssh-client && mkdir /jwt && ssh-keygen -t rsa -f /jwt/jwt_key.rsa -N "" -m pem && ssh-keygen -f /jwt/jwt_key.rsa.pub -e -m pkcs8 > /jwt/jwt_key.rsa.pkcs8

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w -X main.mode=${MODE:-cloudrun}'

FROM scratch as runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/main /app/main
COPY --from=builder /jwt /jwt/

COPY ./templates /templates/

ENTRYPOINT ["/app/main"]
