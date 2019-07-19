FROM golang:1.12

RUN apt update && apt install -yqq protobuf-compiler rsync

WORKDIR /go/src/github.com/zzh8829/kuci
COPY go.mod go.sum ./
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo ./main.go

CMD ["./kuci"]
