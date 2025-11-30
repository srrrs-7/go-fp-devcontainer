FROM golang:1.25 AS builder

WORKDIR /build

# Copy go.mod
COPY go.mod ./

# Copy source files
COPY main.go ./

# Build WebAssembly binary
RUN GOOS=js GOARCH=wasm go build -o main.wasm .

# Copy wasm_exec.js from Go installation
RUN cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./

# Production stage - use nginx to serve static files
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

# Copy static files
COPY static/ ./

# Copy WASM binary and JS glue from builder
COPY --from=builder /build/main.wasm ./
COPY --from=builder /build/wasm_exec.js ./

# Configure nginx for SPA
RUN echo 'server { \
    listen 3000; \
    root /usr/share/nginx/html; \
    index index.html; \
    location / { \
        try_files $uri $uri/ /index.html; \
    } \
    location ~* \.(wasm)$ { \
        add_header Content-Type application/wasm; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
