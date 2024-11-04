FROM golang:1.23 AS builder

COPY ./internal /internal
WORKDIR /internal
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /bin/internal
RUN chmod +x /bin/internal

COPY ./external /external
WORKDIR /external
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /bin/external
RUN chmod +x /bin/external

COPY ./entrypoint.sh /bin/entrypoint
RUN chmod +x /bin/entrypoint

ENTRYPOINT ["entrypoint"]
