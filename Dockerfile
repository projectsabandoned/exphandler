FROM golang:1.19.0 AS builder

WORKDIR /go/src
COPY  . .
RUN go build -o /output/main cmd/main.go

FROM ubuntu

COPY --from=builder /output/main /main

CMD [ "./main" ]
