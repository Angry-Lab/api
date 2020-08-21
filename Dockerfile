FROM golang:1.14-alpine as builder

COPY . /src
WORKDIR /src

ENV CGO_ENABLED 0
ENV GOPRIVATE github.com/Angry-Lab/*

RUN go build -v -o /api ./cmd/api/main.go

FROM alpine:3.11

EXPOSE 6363

WORKDIR /app

RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /api /app/api

ENTRYPOINT ["/app/api"]
