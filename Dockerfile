
FROM golang:1.23.0 as builder

WORKDIR /build

ADD ./ /build

RUN cd /build/ && \
    CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -ldflags '-w -s' -o noah


FROM airdb/base:latest

WORKDIR /app/

COPY --from=builder /build/noah /app/
COPY --from=builder /build/config.yaml /app/

EXPOSE 8080

#ENTRYPOINT ["sleep", "3600"]
#ENTRYPOINT ["/app/apigw"]
CMD ["/app/noah", "run"]