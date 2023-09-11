FROM golang:1.17 as builder

WORKDIR /src

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

FROM alpine:3

RUN apk add --no-cache ca-certificates

COPY --from=builder /src/server /server

ENTRYPOINT ["/server"]
