FROM golang:1.23.0

COPY . . 

RUN go mod download

RUN go build -o application . 

CMD ["application"]

