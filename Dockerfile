# builder image
FROM golang:1.15-buster as builder
ENV SRCDIR /go/src/github.com/skbki/go-http-logger
RUN mkdir -p $SRCDIR
ADD . $SRCDIR
WORKDIR $SRCDIR
RUN mkdir /build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/go-http-logger ./cmd
# runtime image
FROM gcr.io/distroless/static
COPY --from=builder /build/go-http-logger .
ENTRYPOINT [ "./go-http-logger" ]