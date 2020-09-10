FROM golang:alpine

WORKDIR /go/src/app 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]


