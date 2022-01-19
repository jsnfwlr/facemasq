#ARG PCAPV="1.10.1"

FROM golang:1.17-alpine  as godeps
USER root
RUN mkdir -p /build/dist/data /build/dist/api
COPY ./api /build
COPY ./upgrades /build/dist/upgrades
WORKDIR /build
RUN apk add -U --no-cache build-base wget ca-certificates tzdata flex bison linux-headers
RUN go env
RUN wget https://www.tcpdump.org/release/libpcap-1.10.1.tar.gz
RUN tar xvf libpcap-1.10.1.tar.gz && cd libpcap-1.10.1 && ./configure --with-pcap=linux && make
RUN go mod tidy
RUN CGO_ENABLED=1 LD_LIBRARY_PATH="-L/build/libpcap-1.10.1" CGO_LDFLAGS="-L/build/libpcap-1.10.1" CGO_CPPFLAGS="-I/build/libpcap-1.10.1" go build -ldflags '-linkmode external -extldflags -static' --tags "linux sqlite_foreign_keys=1" -o /build/dist/api/facemasq ./
RUN chmod +x /build/dist/api/facemasq

FROM node:16 as vuedeps
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
RUN mkdir -p /build/dist/ui /build/ui
COPY ./ui /build/ui
WORKDIR /build/ui
RUN pnpm install --frozen-lockfile
RUN pnpm build

FROM scratch
WORKDIR /app/data
WORKDIR /app/config
COPY --from=godeps /build/dist /app
COPY --from=vuedeps /build/dist/ui /app/ui
COPY --from=godeps /etc/ssl/certs /etc/ssl/certs
COPY --from=godeps /usr/share/zoneinfo /usr/share/zoneinfo
COPY ./upgrades /app/upgrades

ENV TZ=Etc/UTC
WORKDIR /app/api
CMD ["/app/api/facemasq"]
