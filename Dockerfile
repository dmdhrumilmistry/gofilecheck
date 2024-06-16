FROM cgr.dev/chainguard/go AS builder
WORKDIR /app
COPY . /app
RUN go build -o bin/app cmd/gofilecheck/main.go

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/bin/app /usr/bin/
ENTRYPOINT ["/usr/bin/app"]
