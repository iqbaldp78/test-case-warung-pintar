FROM golang:1.12-alpine as builder

WORKDIR /test-warungpintar
ADD . /test-warungpintar

ENV CGO_ENABLED=0
RUN go build -mod=vendor -tags=jsoniter -o main

FROM alpine:3.8
RUN apk add --no-cache tzdata ca-certificates 
COPY --from=builder /test-warungpintar /app/
WORKDIR /app
CMD ["./main"]