FROM golang AS builder
RUN mkdir -p /go/src/github.com/keithkfield/pg_api/ 
WORKDIR /go/src/github.com/keithkfield/pg_api/
COPY . .
VOLUME /go/src/github.com/keithkfield/pg_api/
RUN go get -d -v github.com/lib/pq
RUN go build main.go



FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/keithkfield/pg_api/main  .
CMD ["./main"]  