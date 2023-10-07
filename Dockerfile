FROM airdb/builder:alpine-go1.20 as builder

WORKDIR /build

ADD ./ /build

RUN cd /build/ && \
	CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -ldflags '-w -s' -o apigw


FROM airdb/base:latest

WORKDIR /app/

COPY --from=builder /build/apigw /app/
COPY --from=builder /build/config.yaml /app/

EXPOSE 8080

#ENTRYPOINT ["sleep", "3600"]
#ENTRYPOINT ["/app/apigw"]
CMD ["/app/apigw"]
