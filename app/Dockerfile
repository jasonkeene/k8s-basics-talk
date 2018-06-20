FROM golang:latest as builder

COPY /main.go /go/main.go

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build \
    -a \
    -installsuffix nocgo \
    -o /app \
    /go/main.go

FROM alpine:latest

ENV ADDR=:80

EXPOSE 80

COPY --from=builder /app /srv/
WORKDIR /srv
CMD [ "/srv/app" ]
