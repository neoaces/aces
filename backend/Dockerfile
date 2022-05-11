#build stage
FROM golang:1.17.9-alpine AS builder
RUN apk add --no-cache git build-base

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build -o /go/bin/app -v

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/app /app
ENTRYPOINT /app

LABEL Name=aced Version=0.0.1

EXPOSE 7979
