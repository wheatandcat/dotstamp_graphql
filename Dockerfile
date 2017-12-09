FROM golang:1.9

COPY . /go/src/app

WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure

EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"] 
