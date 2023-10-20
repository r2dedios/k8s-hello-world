####################
## STEP 1: Build  ##
####################
# Image Base
FROM golang:1.21.3 AS builder

# Code copy
WORKDIR /app
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
RUN go build -o hello_pod ./src





####################
## STEP 2: Run    ##
####################
FROM registry.access.redhat.com/ubi8/ubi:8.8

COPY --from=builder /app/hello_pod /hello_pod

EXPOSE 8080

CMD ["/hello_pod"]

