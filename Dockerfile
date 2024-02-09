FROM golang:1.21-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/AltTube-Go .

# deployment image
FROM alpine:3.5
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/AltTube-Go ./

CMD [ "./AltTube-Go" ]

EXPOSE 8080
