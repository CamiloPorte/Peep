# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /peep
COPY . .

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build -o /docker-peep cmd/main.go
##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build  /docker-peep  /docker-peep

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-peep"]