
FROM caddy:2-builder-alpine

ARG GITHUB_SHA=main
ARG GITHUB_REPOSITORY=caddy-dns/cloudflare

ENV GITHUB_SHA=$GITHUB_SHA
ENV GITHUB_REPOSITORY=$GITHUB_REPOSITORY

ADD . .

RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest && \
    xcaddy build \
    --with "github.com/caddy-dns/cloudflare=."
