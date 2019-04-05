# Build the manager binary
FROM golang:alpine as builder

# Copy in the go src
WORKDIR /go/src/github.com/cloud104/tks-uptimerobot-controller
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager github.com/cloud104/tks-uptimerobot-controller/cmd/manager

# Copy the controller-manager into a thin image
FROM drone/ca-certs 
WORKDIR /
COPY --from=builder /go/src/github.com/cloud104/tks-uptimerobot-controller/manager .
ENTRYPOINT ["/manager"]
