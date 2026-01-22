FROM cgr.dev/chainguard/go:latest AS builder

# SHELL ["/bin/busybox", "sh", "-c"]

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build && \
    mkdir -p /go/bin && \
    mv -v ctop /go/bin/

FROM scratch
ENV TERM=linux
COPY --from=builder /go/bin/ctop /ctop
ENTRYPOINT ["/ctop"]
