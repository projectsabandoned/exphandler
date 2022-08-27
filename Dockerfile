FROM golang:1.19.9 AS builder

WORKDIR /go/src
COPY  . .
RUN go build -o /output/main cmd/main.go

FROM ubuntu:bionic-20171114

COPY --from=builder /output/main /main

CMD [ "./main" ]
