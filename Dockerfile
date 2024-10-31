FROM golang:1.22.3

WORKDIR /go/src/GO-API
COPY . .

RUN go build -o bin/server cmd/main.go
CMD [ "./bin/server" ]