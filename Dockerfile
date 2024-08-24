FROM golang:1.23.0-bookworm AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . . 

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o application -a -ldflags="-s -w" -installsuffix cgo

RUN upx --ultra-brute -qq application && upx -t application 

FROM scratch

COPY --from=build /app/application /application

ENTRYPOINT ["/application"]

