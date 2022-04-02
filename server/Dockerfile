FROM hub.deepin.com/library/golang:bullseye AS builder
WORKDIR /src
COPY . .
RUN go build -o main main.go

FROM hub.deepin.com/public/uniteos:2021
ENV ALCHEMY_FURNACE_ROOT=/service
WORKDIR /service
RUN set -eu; \
    apt-get update; \
    apt-get install -y openssh-client
COPY --from=builder /src/main /usr/bin/alchemy-furnace
ENTRYPOINT ["alchemy-furnace"]