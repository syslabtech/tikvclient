FROM golang:1.23.0-bookworm AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . . 

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server -a -ldflags="-s -w" -installsuffix cgo

RUN upx --ultra-brute -qq server && upx -t server 

FROM scratch

COPY --from=build /app/server /server

ENTRYPOINT ["/server"]

