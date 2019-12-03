FROM golang:alpine3.10
MAINTAINER Mohammed Aadil "mailtoaadilhanif@gmail.com"

RUN apk update && apk add git build-base python-dev py-pip 
RUN go get -u github.com/golang/dep/cmd/dep
COPY . /go/src/github.com/Mohammed-Aadil/ds-storage
COPY ./Gopkg.toml /go/src/github.com/Mohammed-Aadil/ds-storage

WORKDIR /go/src/github.com/Mohammed-Aadil/ds-storage
RUN dep ensure
RUN go test cmd/main.go -v
RUN go build cmd/main.go
CMD [ "./main" ]
EXPOSE 8001
