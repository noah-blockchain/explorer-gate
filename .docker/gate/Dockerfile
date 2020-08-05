FROM golang:1.13.7-alpine as builder

ARG VERSION

RUN apk add --no-cache make gcc musl-dev linux-headers
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -ldflags="-X 'github.com/noah-blockchain/explorer-gate/v2/src/api.Version=$VERSION' -X 'main.Version=$VERSION'" -o ./builds/linux/gate ./cmd/gate.go

FROM alpine:3.7

COPY --from=builder /app/builds/linux/gate /usr/bin/gate
RUN addgroup noahuser && adduser -D -h /noah -G noahuser noahuser
USER noahuser
WORKDIR /noah
CMD ["/usr/bin/gate"]

