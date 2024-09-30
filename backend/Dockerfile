FROM golang:1.23.1-alpine

RUN mkdir /go/src/backend
RUN apk update && apk add git

WORKDIR /go/src/backend

ADD . /go/src/backend

RUN go mod download

RUN go install github.com/air-verse/air@v1.60.0

CMD ["air", "-c", ".air.toml"]
