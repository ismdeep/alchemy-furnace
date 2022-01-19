FROM hub.deepin.com/library/golang:alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o main main.go

FROM hub.deepin.com/library/alpine:latest
WORKDIR /service
COPY --from=builder /src/main /usr/bin/alchemy-furnace
ENTRYPOINT ["alchemy-furnace"]