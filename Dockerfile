FROM golang:1.12

RUN apt update && apt install -yqq \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg2 \
    software-properties-common

RUN curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
RUN sudo add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/debian \
    $(lsb_release -cs) \
    stable"
RUN apt-get update && apt-get install -yqq \
    docker-ce-cli

RUN curl -L https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

WORKDIR /kuci

COPY go.mod go.sum ./
RUN go mod vendor

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo

CMD ["./kuci"]
