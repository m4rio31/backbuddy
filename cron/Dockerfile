# Stage 1: Compile GO app
FROM golang:1.24 AS builder

WORKDIR /go/src/go.code.github.com

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o cron ./cron


# Stage 2: Setup image for cronjob
FROM alpine:3.22 AS runner

COPY --from=builder /go/src/go.code.github.com/cron /usr/local/bin/

COPY ./daemon/cron/crontab.txt /crontab.txt

RUN /usr/bin/crontab /crontab.txt

CMD ["/usr/sbin/crond", "-f"]