FROM golang:1.21-alpine

RUN apk add --no-cache upx gcc musl-dev ca-certificates
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=1 go build -o curl --ldflags '-linkmode external -extldflags "-static"' . && strip /app/curl && upx -9 /app/curl

FROM scratch
COPY passwd /etc/passwd
COPY --from=0 /app/curl /curl
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER nobody

ENTRYPOINT ["/curl"]