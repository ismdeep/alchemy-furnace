FROM golang:bullseye AS server-builder
WORKDIR /src
COPY ./server .
RUN set -eux; \
    go mod tidy; \
    go mod download; \
    go mod vendor; \
    go build -mod vendor -o main main.go

FROM node:16 AS web-builder
WORKDIR /src
COPY ./web .
RUN yarn install
RUN yarn build

# main
FROM debian:bullseye
ENV ALCHEMY_FURNACE_ROOT=/service
WORKDIR /service
RUN set -eux; \
    apt-get update; \
    apt-get install -y openssh-client supervisor nginx
COPY --from=server-builder /src/main /usr/bin/alchemy-furnace
COPY --from=web-builder /src/dist /usr/share/nginx/html/
COPY ./.data/nginx/default /etc/nginx/sites-available/default
COPY ./.data/supervisord.conf /etc/supervisor/supervisord.conf
COPY ./.data/config.yaml /service/config.yaml
EXPOSE 80
ENTRYPOINT ["/usr/bin/supervisord"]