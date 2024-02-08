FROM alpine:3.12
WORKDIR /opt
COPY .env .
COPY ustaz-service-go .

CMD ["/opt/ustaz-service-go"]