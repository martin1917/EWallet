FROM golang:1.19-alpine3.15 as builder
WORKDIR /go/src/EWallet
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./build/app ./cmd/main.go


FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /project
COPY --from=builder /go/src/EWallet/build/app ./
COPY ./.env ./
EXPOSE 8080
CMD ["./app"]