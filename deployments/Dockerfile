FROM golang:1.21.1-alpine3.18 as builder

COPY . /go/src/dike

WORKDIR /go/src/dike

RUN go mod tidy

RUN go build cmd/main.go


FROM alpine:3.18.0

WORKDIR /root

COPY --from=builder /go/src/dike/dike .

RUN chmod +x dike

ENV GIN_MODE=release

CMD ["./dike"]