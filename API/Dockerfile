FROM golang:1.15.6-alpine3.12

RUN apk --no-cache add gcc g++ make git

LABEL Maintainer="Vague Coder <vaguecoder0to.n@gami.com>"

WORKDIR $GOPATH/src/github.com/VagueCoder/RESTful-API-Go-Endpoint/API

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["RESTful-API-Go-Endpoint"]