FROM cgr.dev/chainguard/go:latest-dev as builder
WORKDIR app
COPY . /app/
RUN CGO_ENABLED=0 go build -o gupi .

FROM cgr.dev/chainguard/static
COPY --from=builder /app/gupi /gupi
ENTRYPOINT ["/gupi"]
