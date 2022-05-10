# syntax=docker/dockerfile:1
#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
# Get the required GOLANG dependencys
RUN go mod download
RUN go build -o /go/bin/app

#final stage // reduces the file size of the container
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy over the images
COPY --from=builder /go/bin/app /app
CMD [ "/app" ]

LABEL Name=aces Version=0.0.1
EXPOSE 8080
