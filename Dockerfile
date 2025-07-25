FROM cgr.dev/chainguard/go:latest AS builder

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .
RUN make build && \
    mkdir -p /go/bin && \
    mv -v ctop /go/bin/

FROM scratch
ENV TERM=linux
COPY --from=0 /go/bin/ctop /ctop
ENTRYPOINT ["/ctop"]
