FROM golang:latest AS builder

RUN apt-get update

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ADD . /ArquicturaSW

WORKDIR /ArquicturaSW
RUN go mod init ArquicturaSW
RUN go mod tidy
RUN go build -o ArquicturaSW .
RUN chmod +x /ArquicturaSW

ENTRYPOINT [" ./main"]
