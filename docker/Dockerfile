# syntax=docker/dockerfile:1

FROM golang:1.19 AS godeps
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS
RUN apt-get update && apt-get install -y binutils gcc g++ pkg-config make wget ca-certificates tzdata flex bison libpcap-dev
RUN mkdir -p /dist/api
WORKDIR /build
COPY ./api /build
COPY ./upgrades /dist/upgrades
COPY ./templates /dist/templates
RUN CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" --tags "linux sqlite_foreign_keys=1" -o /dist/api/facemasq .
RUN chmod +x /dist/api/facemasq

FROM node:16 as vuedeps
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
RUN mkdir -p /build/dist/web /build/web
COPY ./web /build/web
WORKDIR /build/web
RUN pnpm install --frozen-lockfile
RUN pnpm list
RUN pnpm build

FROM scratch
WORKDIR /app/data
WORKDIR /app/config
COPY --from=godeps /dist /app
COPY --from=vuedeps /build/dist/web /app/web
COPY --from=godeps /etc/ssl/certs /etc/ssl/certs
COPY --from=godeps /usr/share/zoneinfo /usr/share/zoneinfo
COPY ./upgrades /app/upgrades
COPY ./templates /app/templates

ENV TZ=Etc/UTC
WORKDIR /app/api
CMD ["/app/api/facemasq"]
