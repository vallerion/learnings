FROM golang:1.19 AS build

WORKDIR /app

COPY ../.. ./

RUN CGO_ENABLED=0 go build -o /app/bin main.go

FROM alpine as run

WORKDIR /

COPY --from=build /app/bin /app/bin

ARG VER_TAG
ENV VER_TAG=$VER_TAG

EXPOSE 8081

ENTRYPOINT ["/app/bin"]