FROM hub.deepin.com/library/golang:bullseye AS builder
WORKDIR /src
COPY . .
RUN go build -o main main.go

FROM hub.deepin.com/public/uniteos:2021
WORKDIR /service
COPY ./entrypoint.bash /usr/bin/entrypoint.bash
RUN set -eu; \
    curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun; \
    apt-get update; \
    apt-get install -y gcc g++ make cmake git
COPY --from=builder /src/main /usr/bin/alchemy-furnace
ENTRYPOINT ["entrypoint.bash"]