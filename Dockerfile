FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/EORUG/avitotest/
WORKDIR /go/src/github.com/EORUG/avitotest
RUN go mod download
COPY . /go/src/github.com/EORUG/avitotest
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/avitotest github.com/EORUG/avitotest

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/EORUG/avitotest/build/avitotest /usr/bin/avitotest
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/bucketeer"]