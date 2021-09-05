FROM golang:1.16 as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN make build

FROM alpine:3.13
RUN apk add --no-cache curl
RUN mkdir -p /app/bin
WORKDIR /app
RUN mkdir -p tmp
COPY --from=builder /build/bin /usr/local/bin/

# HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD [ "sh", "-c", "curl --user-agent 'Internal Container HealthCheck' -H 'Content-Type: text/plain' -f http://localhost:3000/health || exit 1;" ]

CMD ["server"]
