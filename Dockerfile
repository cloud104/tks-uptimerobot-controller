#### Kubebuilder
FROM alpine as kubebuilder
ENV VERSION=1.0.8 \
    ARCH=amd64
RUN apk add --update --no-cache curl
RUN curl -L "https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${VERSION}/kubebuilder_${VERSION}_linux_${ARCH}.tar.gz" -o kubebuilder.tar.gz && \
    tar -zxvf kubebuilder.tar.gz && \
    mv "kubebuilder_${VERSION}_linux_${ARCH}" /usr/local/kubebuilder

#### Dep
FROM golang:alpine as dep
RUN apk add --update --no-cache git curl && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

#### Dev Layer
FROM golang:alpine as dev
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk add --update --no-cache git
COPY --from=dep /go/bin/dep /go/bin/dep
COPY --from=kubebuilder /usr/local/kubebuilder /usr/local/kubebuilder
WORKDIR $GOPATH/src/github.com/cloud104/tks-uptimerobot-controller
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

#### Builder
FROM dev as builder
COPY pkg/ pkg/
COPY cmd/ cmd/
RUN go build -o /manager github.com/cloud104/tks-uptimerobot-controller/cmd

#### Prod
#FROM drone/ca-certs
FROM alpine
RUN apk add -U --no-cache ca-certificates
ENV HOME=/root
COPY --from=builder /manager /manager
ENTRYPOINT ["/manager"]
