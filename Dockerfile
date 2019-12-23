FROM golang:1.8


WORKDIR /go/src/go-web-app
COPY . .


RUN go get -d ./...
RUN go install -v ./...

CMD ["go-web-app"]
