FROM golang:alpine

RUN apk add alpine-sdk

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/rianby64/gin-example
COPY . .

RUN dep ensure -v

RUN go build

CMD [ "./gin-example" ]

#CMD [ "dlv", "debug", "--headless", "--listen=:2345", "--log", "--api-version", "2", "--accept-multiclient", "github.com/rianby64/gin-example" ]