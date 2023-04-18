FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/user

# Install dependencies
RUN apk --update --no-cache add gcc libtool make musl-dev protoc git

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN make init && go env -w GOPROXY=https://proxy.golang.com.cn,direct
COPY . .
RUN make proto tidy build

# Deployment container
FROM scratch

COPY --from=builder /go/src/user/user /user
ENTRYPOINT ["/post"]
CMD []
