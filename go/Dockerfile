FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY ./app.go .

RUN go build app.go

FROM golang:alpine AS application

WORKDIR /go/src/app

COPY --from=builder /go/src/app/app .

CMD ["./app"]

EXPOSE 8000
