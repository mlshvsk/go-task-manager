FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/pressly/goose/cmd/goose
RUN go install -v github.com/pressly/goose/cmd/goose
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]