FROM golang as build
ADD . /usr/local/go/src/giligili
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /usr/local/go/src/giligili
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7
ENV MYSQL_DSN=""
ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB="0"
ENV GIN_MODE="release"
ENV PORT=3000

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /usr/local/go/src/giligili/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]