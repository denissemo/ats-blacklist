FROM golang:1.15

RUN export GOPATH=/go

RUN go env GOPATH

COPY . /go/src/ats-blacklist

WORKDIR /go/src/ats-blacklist

RUN go build -o ./build -v .

EXPOSE 3000

CMD ./build/ats-blacklist