FROM golang:1.14-alpine AS compile
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 \
    go build -a -o app .

FROM alpine:3.11
COPY --from=compile /app/app /
ENV GIN_BIND=[::]:80\
    GIN_MODE=release

ENTRYPOINT ["/app"]
