FROM golang:1.22 AS build-stage

WORKDIR /

COPY . .
RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /monkey .

# multi-stage build to keep the image minimal
FROM docker.io/alpine:edge AS release-stage

WORKDIR /

COPY --from=build-stage /monkey /monkey

USER root:root

ENTRYPOINT ["/monkey"]
