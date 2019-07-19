FROM golang:1.12

RUN apt update && apt install -yqq protobuf-compiler rsync

WORKDIR /kuci

COPY go.mod go.sum ./
RUN go mod vendor

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo

CMD ["./kuci"]
