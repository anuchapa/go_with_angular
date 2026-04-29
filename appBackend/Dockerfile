FROM golang:1.25 as build 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:3.20
WORKDIR /app

COPY --from=build /app/app .

CMD ["./app"]