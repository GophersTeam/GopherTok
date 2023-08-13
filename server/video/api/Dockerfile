FROM golang:1.20.5-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY server/video/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/video server/video/api/video.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

ENV TZ Asia/Shanghai
ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

WORKDIR /app
COPY --from=builder /app/video /app/video
COPY --from=builder /app/etc /app/etc

CMD ["./video", "-f", "etc/video.yaml"]