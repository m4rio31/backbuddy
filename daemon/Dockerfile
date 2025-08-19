# Stage 1: Compile GO app
FROM golang:1.24 AS builder

WORKDIR /go/src/go.code.github.com

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o daemon ./daemon


# Stage 2: Setup final image
FROM alpine:3.17

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk add --no-cache cronie

COPY --from=builder --chown=appuser:appgroup /go/src/go.code.github.com/daemon /usr/local/bin/

COPY --chown=appuser:appgroup daemon/cron/entry.sh /usr/local/bin/entry.sh
RUN chmod +x /usr/local/bin/entry.sh

COPY --chown=appuser:appgroup daemon/cron/crontab.txt /etc/crontabs/appuser

USER appuser

CMD ["/usr/local/bin/entry.sh"]