####################
## STEP 1: Build  ##
####################
# Image Base
FROM golang:1.18 AS builder

# Code copy
RUN mkdir /go/hello_world
WORKDIR /go/hello_world
COPY . .

# Environment variable to Static-link compilation
ENV CGO_ENABLED 0
# Work Go Path
ENV GOPATH /go
# GO OS system for compatibility compilation
ENV GOOS linux
# GO OS system architecture
ENV GOARCH amd64

# Run command
RUN go build -o hello_world main.go





####################
## STEP 2: Run    ##
####################
FROM scratch

# Binary
COPY --from=builder /go/hello_world/hello_world /hello_world

EXPOSE 8080

ENTRYPOINT ["/hello_world"]

