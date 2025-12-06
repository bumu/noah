FROM golang:1.25-alpine AS builder

WORKDIR /build

ADD ./ /build

RUN cd /build/ && \
    CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -ldflags '-w -s' -o noah cmd/noah/main.go


FROM alpine:3.23

ENV AIRDB_RUNNING_CONTEXT=docker

RUN apk add --no-cache --virtual .build-deps \
    tzdata tar gzip \
    curl netcat-openbsd \
    bash bash-completion ca-certificates

WORKDIR /opt/app

COPY --from=builder /build/noah /opt/app//
COPY --from=builder /build/configs/config.yaml /opt/app/

EXPOSE 8080

#ENTRYPOINT ["sleep", "3600"]
#ENTRYPOINT ["/app/apigw"]
CMD ["/opt/app/noah", "run"]
