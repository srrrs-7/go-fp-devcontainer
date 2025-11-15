# Builder
FROM golang:1.25 AS builder
ADD ./workspace/main/apps /apps
WORKDIR /apps/api/src
ENV GOOS=linux
ENV GO_ARCH=amd64
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /usr/local/bin/api ./cmd

# Runner
FROM alpine:latest AS runner
# Copy built application and dependencies
COPY --from=builder /usr/local/bin/api /usr/local/bin/api
EXPOSE 8080
USER api
CMD ["/usr/local/bin/api"]