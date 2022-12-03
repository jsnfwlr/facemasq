# syntax=docker/dockerfile:1

FROM golang:1.19 AS godeps
RUN go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
WORKDIR /code
VOLUME /code
VOLUME /output

CMD ["gomarkdoc",  "--output={{.Dir}}/README.md", "/code/..."]