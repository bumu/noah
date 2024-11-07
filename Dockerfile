
FROM golang:1.23 AS builder

WORKDIR /build

ADD ./ /build

RUN cd /build/ && \
    CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -ldflags '-w -s' -o noah cmd/noah/main.go


FROM airdb/alpine:latest

WORKDIR /app/

COPY --from=builder /build/noah /app/
COPY --from=builder /build/configs/config.yaml /app/

EXPOSE 8080

#ENTRYPOINT ["sleep", "3600"]
#ENTRYPOINT ["/app/apigw"]
CMD ["/app/noah", "run"]
