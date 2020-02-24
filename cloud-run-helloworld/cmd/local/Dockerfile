FROM golang:1.13-alpine as builder
RUN apk add --no-cache git
COPY ./ /root/app
WORKDIR /root/app
RUN go build -o app ./cmd/local

FROM alpine:3.10
RUN apk add --no-cache ca-certificates
COPY --from=builder /root/app/app /usr/local/bin
EXPOSE 8080
ENTRYPOINT "app"
