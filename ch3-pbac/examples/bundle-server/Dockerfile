# FROM golang:1.19.2-alpine3.16 as builder
FROM golang:1.19.2 as builder
# RUN apk add git
# RUN apk add git build-base
RUN apt install git
WORKDIR /
COPY . ./
# Disable default GOPROXY
RUN go env -w GOPROXY=direct
# RUN GOOS=linux GOARCH=amd4 go build -a -o main.bin .
RUN GOOS=linux GOARCH=amd64 go build -o main.bin .

# FROM scratch
# FROM alpine:3.16
FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder main.bin main
EXPOSE 8080
ENTRYPOINT ["/main"]
