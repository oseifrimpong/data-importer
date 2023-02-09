# --- Build stage ---
FROM golang:1.19 AS builder
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
# Copy the dependency definition
COPY ./go.mod .
COPY ./go.sum .
# Download dependencies
RUN go mod download
# Copy the source code
COPY . .
# Build for release
RUN CGO_ENABLED=0 GOOS=linux go build 

# --- Final stage ---
FROM alpine:3

COPY --from=builder /app/ohlc-data-api /
COPY ./deploy/scripts/app/init.sh ./deploy/scripts/app/entrypoint.sh /

RUN chmod +x /init.sh && chmod +x /entrypoint.sh
ENTRYPOINT ["/init.sh", "/entrypoint.sh"]

EXPOSE 3000
LABEL Name=ohlc-data-api

CMD /ohlc-data-api
