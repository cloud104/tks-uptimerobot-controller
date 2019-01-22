# Go envs
FROM golang:alpine AS dev
RUN apk add --update --no-cache git curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ENV APP_HOME $GOPATH/src/github.com/cloud104/tks-uptimerobot-controller
WORKDIR $APP_HOME
COPY .kube/config Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
RUN ls $APP_HOME
VOLUME ["$APP_HOME"]
# Builder
FROM dev AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint *.go

# Prod
FROM drone/ca-certs
COPY --from=builder /entrypoint /usr/local/bin/entrypoint
ENTRYPOINT ["entrypoint"]