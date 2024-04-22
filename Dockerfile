FROM golang:1.20 as builder

# build as a non-root user
RUN addgroup --gid 10001 --system nonroot && adduser -u 10000 --system --gid 10001 --home /home/nonroot nonroot
RUN chown -R nonroot:nonroot /home/nonroot
USER nonroot

WORKDIR /code

COPY --chown=nonroot:nonroot . /code

RUN go mod download
RUN mkdir /code/build \
 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /code/build/bot /code/cmd/bot/bot.go


# lighter runtime
FROM alpine:latest

ENV UNCIVBOT_TOKEN=""
ENV UNCIVBOT_DEBUG_LOGS=0
ENV UNCIVBOT_POOL_SIZE=4
ENV UNCIVBOT_DB_PATH="/db/db.sqlite"

# setup the non-root user
RUN addgroup --gid 10001 --system nonroot && adduser -u 10000 --system -G nonroot --home /home/nonroot nonroot
RUN chown -R nonroot:nonroot /home/nonroot

RUN apk --no-cache add ca-certificates \
    && update-ca-certificates
VOLUME /db
WORKDIR /app
COPY --from=builder --chown=nonroot:nonroot /code/build/bot /app/bot

USER nonroot
CMD ["/app/bot"]  
