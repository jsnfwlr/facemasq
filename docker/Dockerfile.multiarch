# syntax=docker/dockerfile:1
ARG PLATFORMS="linux/386 linux/amd64 linux/arm64"
# linux/arm/v5 linux/arm/v6 linux/arm/v7 linux/mips linux/mipsle linux/mips64 linux/mips64le linux/ppc64le linux/riscv64 linux/s390x windows/386 windows/amd64
FROM --platform=$BUILDPLATFORM crazymax/goxx:1.19 AS godeps

ARG PLATFORMS
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS
# make wget
RUN --mount=type=cache,sharing=private,target=/var/cache/apt --mount=type=cache,sharing=private,target=/var/lib/apt/lists goxx-apt-get install -y ca-certificates tzdata flex bison libpcap-dev
RUN mkdir -p /build/api /build/i18n /dist/api /dist/logs /dist/extensions
RUN <<EOT
export GOXX_SKIP_APT_PORTS=1
goxx-apt-get update
for p in $PLATFORMS; do
  TARGETPLATFORM=$p goxx-apt-get install -y binutils gcc g++ pkg-config
done
EOT
WORKDIR /build/api
COPY ./i18n /build/i18n
COPY ./api /build/api
COPY ./upgrades /dist/upgrades
COPY ./templates /dist/templates
RUN --mount=type=cache,target=/root/.cache --mount=type=cache,target=/go/pkg/mod CGO_ENABLED=1 goxx-go build --tags "linux sqlite_foreign_keys=1" -ldflags "-linkmode external -extldflags -static" -o /dist/api/facemasq .
RUN --mount=type=cache,target=/root/.cache --mount=type=cache,target=/go/pkg/mod find ./extensions -maxdepth 1 -type d | grep "\/.*\/" | xargs -n 1 -I {} bash -c 'CGO_ENABLED=1 goxx-go build --tags "linux" -buildmode=plugin {}'
RUN mv *.so /dist/extensions
RUN chmod +x /dist/api/facemasq
RUN chmod +x /dist/extensions/*

FROM --platform=$BUILDPLATFORM node:16 as vuedeps
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
RUN mkdir -p /build/dist/web /build/web /buld/i18n
COPY ./i18n /build/i18n
COPY ./web /build/web
WORKDIR /build/web
RUN pnpm install --frozen-lockfile
RUN pnpm build

FROM scratch
COPY --from=godeps /dist /app
COPY --from=vuedeps /build/dist/web /app/web
COPY --from=godeps /etc/ssl/certs /etc/ssl/certs
COPY --from=godeps /usr/share/zoneinfo /usr/share/zoneinfo
COPY ./upgrades /app/upgrades
COPY ./templates /app/templates
COPY ./i18n /app/i18n

ENV TZ=Etc/UTC
VOLUME /data
VOLUME /export
WORKDIR /app/api
CMD ["/app/api/facemasq"]
