FROM golang:1.18-buster AS builder

ENV GO111MODULE on

COPY . /usr/local/src/repo
WORKDIR /usr/local/src/repo

# RUN go mod download github.com/x-asia/kauche-personal-shota-kohno/application/echo

RUN go build -o server

# FROM arm64v8/alpine:latest
FROM gcr.io/distroless/base:3c213222937de49881c57c476e64138a7809dc54
EXPOSE 8080
COPY --from=builder /usr/local/src/repo/server /usr/local/bin/server
ENTRYPOINT ["/usr/local/bin/server"]
