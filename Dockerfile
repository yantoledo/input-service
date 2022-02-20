FROM golang:1.17.7-alpine3.15 as builder

WORKDIR /go/src

COPY . . 
RUN go mod download
RUN go build -o . -mod=readonly ./main.go

EXPOSE 5003

CMD ["go", "run", "./main.go"]