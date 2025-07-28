# -------- build stage --------
FROM golang:1.22-bookworm AS builder
WORKDIR /app
COPY . .
RUN cd backend/gameplay-engine && \
    go mod download && \
    CGO_ENABLED=0 go build -o /go/bin/engine ./cmd/engine

# -------- run stage ----------
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /go/bin/engine /engine
ENTRYPOINT ["/engine"]
