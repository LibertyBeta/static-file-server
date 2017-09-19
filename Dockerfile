FROM golang:latest as builder

COPY serve.go /

RUN mkdir /web

COPY blank.pdf /web/file.pdf

WORKDIR /

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o serve .

FROM scratch
COPY --from=builder /serve /
COPY --from=builder /web/file.pdf /web/file.pdf
CMD ["/serve"]
